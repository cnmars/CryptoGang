package init

import (
	"context"
	"flag"
	"fmt"

	// admin "golddigger/app/system/router"
	system "golddigger/app/system/router"
	user "golddigger/app/user/router"

	// "net/http"
	// "golddigger/app/system/apis"
	"golddigger/common/jobs"
	"golddigger/common/wsserver"
	"golddigger/global"
	"golddigger/utils"

	"golddigger/common/etherscan"
	internal "golddigger/init/interobj"
	"golddigger/middleware"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// GormMysql 初始化Mysql数据库

func InitSystemOnline(configPath ...string) {
	//要先初始化日志，不然后续的配置出错没法写日志
	//initViper方法中会将配置文件解析到global.GD_CONFIG变量中
	if len(configPath) == 0 {
		global.GD_VP = initViper()
	} else {
		global.GD_VP = initViper(configPath[0])
	}

	//初始化全局日志
	initZap()

	// 初始化redis服务
	if global.GD_CONFIG.System.UseMultipoint || global.GD_CONFIG.System.UseRedis {
		initRedis()
	}

	//初始化数据库连接
	initDBConnection()
	initRouters()

	// initEthClientConnect()
	initChainProvider()
	initWebsocketServer()
	otherInit()
}

// 执行控制台命令时也需要一些初始化
func InitSystemForCMD(configPath ...string) {

	//假如线上环境在跑的话就已经初始化过了，这里不再初始化

	if global.GD_VP == nil {
		//要先初始化日志，不然后续的配置出错没法写日志
		//initViper方法中会将配置文件解析到global.GD_CONFIG变量中
		if len(configPath) == 0 {
			global.GD_VP = initViper()
		} else {
			global.GD_VP = initViper(configPath[0])
		}
	}
	if global.GD_LOG == nil {
		//初始化全局日志
		initZap()
	}
	if global.GD_REDIS == nil {
		// 初始化redis服务
		if global.GD_CONFIG.System.UseMultipoint || global.GD_CONFIG.System.UseRedis {
			initRedis()
		}
	}

	if global.GD_DB == nil {
		//初始化数据库连接
		initDBConnection()
	}
}

func RunWindowsServer() {

	// 从db加载jwt数据
	// if global.GD_DB != nil {
	// 	system.LoadAll()
	// }

	Router := global.GD_ENGIN
	Router.Static("/form-generator", "./resource/page")
	address := fmt.Sprintf(":%d", global.GD_CONFIG.System.Port)
	s := internal.InitServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.GD_LOG.Info("server run success on ", zap.String("address", address))

	fmt.Printf(``, address)
	global.GD_LOG.Error(s.ListenAndServe().Error())
}

func initDBConnection() {
	m := global.GD_CONFIG.Mysql
	if m.Dbname == "" {
		return
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), internal.Gorm.Config()); err != nil {
		return
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		global.GD_DB = db
	}
}
func initRedis() {
	redisCfg := global.GD_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	//不知道这里为什么要换成无参的
	// pong, err := client.Ping().Result()

	if err != nil {
		fmt.Println(zap.Error(err))
		global.GD_LOG.Error("redis connect ping failed, err:", zap.Error(err))
	} else {
		global.GD_LOG.Info("redis connect ping response:", zap.String("pong", pong))
		global.GD_REDIS = client
	}
}

func initZap() {
	if ok, _ := utils.PathExists(global.GD_CONFIG.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", global.GD_CONFIG.Zap.Director)
		_ = os.Mkdir(global.GD_CONFIG.Zap.Director, os.ModePerm)
	}

	cores := internal.Zap.GetZapCores()
	logger := zap.New(zapcore.NewTee(cores...))

	if global.GD_CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	global.GD_LOG = logger
	// return logger
}

// 优先级: 命令行 > 环境变量 > 默认值
func initViper(path ...string) *viper.Viper {
	var config string

	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file.")
		flag.Parse()
		if config == "" { // 判断命令行参数是否为空
			if configEnv := os.Getenv(global.ConfigEnv); configEnv == "" { // 判断 internal.ConfigEnv 常量存储的环境变量是否为空
				switch gin.Mode() {
				case gin.DebugMode:
					fmt.Println("111")
					config = global.ConfigDefaultFile
					fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.EnvGinMode, global.ConfigDefaultFile)
				case gin.ReleaseMode:
					fmt.Println("222")
					config = global.ConfigReleaseFile
					fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.EnvGinMode, global.ConfigReleaseFile)
				case gin.TestMode:
					fmt.Println("333")
					config = global.ConfigTestFile
					fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.EnvGinMode, global.ConfigTestFile)
				}
			} else { // internal.ConfigEnv 常量存储的环境变量不为空 将值赋值于config
				config = configEnv
				fmt.Printf("您正在使用%s环境变量,config的路径为%s\n", global.ConfigEnv, config)
			}
		} else { // 命令行参数不为空 将值赋值于config
			fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%s\n", config)
		}
	} else { // 函数传递的可变参数的第一个值赋值于config
		config = path[0]
		fmt.Printf("您正在使用func Viper()传递的值,config的路径为%s\n", config)
	}

	v := viper.New()
	// v.SetConfigFile("/home/spike/project/myown/GD/config/config.yaml")
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Viper --> config file changed:", e.Name)
		if err = v.Unmarshal(&global.GD_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&global.GD_CONFIG); err != nil {
		fmt.Println(err)
	}

	// root 适配性 根据root位置去找到对应迁移位置,保证root路径有效
	// global.GD_CONFIG.AutoCode.Root, _ = filepath.Abs("..")
	return v
}

func initRouters() {
	var Router *gin.Engine
	h := global.GD_ENGIN
	if h == nil {
		h = gin.New()
		global.GD_ENGIN = h
	}
	Router = gin.Default()
	// adminRouter := admin.AdminGroup.Example
	// userRouter := user.UserGroup.Common
	// 如果想要不使用nginx代理前端网页，可以修改 web/.env.production 下的
	// VUE_APP_BASE_API = /
	// VUE_APP_BASE_PATH = http://localhost
	// 然后执行打包命令 npm run build。在打开下面4行注释
	// Router.LoadHTMLGlob("./dist/*.html") // npm打包成dist的路径
	// Router.Static("/favicon.ico", "./dist/favicon.ico")
	// Router.Static("/static", "./dist/assets")   // dist里面的静态资源
	// Router.StaticFile("/", "./dist/index.html") // 前端网页入口页面

	// Router.StaticFS(global.GD_CONFIG.Local.Path, http.Dir(global.GD_CONFIG.Local.StorePath)) // 为用户头像和文件提供静态地址
	if global.GD_CONFIG.SslConfig.Enable {
		Router.Use(middleware.TlsHandler()) 
	}
	// 跨域，如需跨域可以打开下面的注释
	// Router.Use(middleware.Cors()) // 直接放行全部跨域请求
	//Router.Use(middleware.CorsByRules()) // 按照配置的规则放行跨域请求
	//global.GD_LOG.Info("use middleware cors")
	// Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// global.GD_LOG.Info("register swagger handler")
	// 方便统一添加路由组前缀 多服务器上线使用

	privateGroup := Router.Group("")
	publicGroup := Router.Group("")
	user.InitUserPublicRouter(publicGroup)
	//基础路由不做鉴权：登录、获取nonce，关于建立websocket建立连接是应该放在哪个环节？
	system.InitSystemPublicBasicRouter(privateGroup)
	//测试的时候不带token，正式环境要带token
	privateGroup.Use(middleware.JWTAuth())
	// privateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	system.InitSystemModuleRouter(privateGroup)

	// {
	// 	// 健康监测
	// 	PublicGroup.GET("/health", func(c *gin.Context) {
	// 		c.JSON(200, "ok")
	// 	})
	// }
	{
		// adminRouter.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
		// adminRouter.InitInitRouter(PublicGroup) // 自动初始化相关
	}
	// PrivateGroup := Router.Group("")
	// privateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		// adminRouter.InitApiRouter(PrivateGroup)                 // 注册功能api路由
		// system.InitJwtRouter(privateGroup)                 // jwt相关路由
		// adminRouter.InitUserRouter(PrivateGroup)                // 注册用户路由
		// adminRouter.InitMenuRouter(PrivateGroup)                // 注册menu路由
		// adminRouter.InitSystemRouter(PrivateGroup)              // system相关路由
		// adminRouter.InitCasbinRouter(PrivateGroup)              // 权限相关路由
		// adminRouter.InitAutoCodeRouter(PrivateGroup)            // 创建自动化代码
		// adminRouter.InitAuthorityRouter(PrivateGroup)           // 注册角色路由
		// adminRouter.InitSysDictionaryRouter(PrivateGroup)       // 字典管理
		// adminRouter.InitAutoCodeHistoryRouter(PrivateGroup)     // 自动化代码历史
		// userRouter.InitSysOperationRecordRouter(PrivateGroup) // 操作记录
		// adminRouter.InitSysDictionaryDetailRouter(PrivateGroup) // 字典详情管理
		// adminRouter.InitAuthorityBtnRouterRouter(PrivateGroup)  // 字典详情管理

		// userRouter.InitExcelRouter(PrivateGroup)                 // 表格导入导出
		// userRouter.InitCustomerRouter(PrivateGroup)              // 客户路由
		// userRouter.InitFileUploadAndDownloadRouter(PrivateGroup) // 文件上传下载功能路由

		// Code generated by golddigger/server Begin; DO NOT EDIT.

		// Code generated by golddigger/server End; DO NOT EDIT.
	}

	// InstallPlugin(Router) // 安装插件

	global.GD_LOG.Info("router register success")
	global.GD_ENGIN = Router
}

//	func initEthClientConnect() {
//		var (
//			ctx         = context.Background()
//			url         = "https://mainnet.infura.io/v3/9a80a6abe921408298d4359a31faa3f0"
//			client, err = ethclient.DialContext(ctx, url)
//		)
//		if err == nil {
//			global.GD_CLIENT = client
//			global.GD_LOG.Info(utils.Green("连接以太坊客户端成功!"))
//		} else {
//			global.GD_LOG.Error(utils.Red("连接客户端出错："))
//			global.GD_LOG.Error(err.Error())
//		}
//	}
func initWebsocketServer() {
	global.GD_WS = &wsserver.PubSub{}
	global.GD_WS.IsStop = false
	global.GD_WS.ThirdPartData = make(chan wsserver.WsMessage, 100)

	//###########独立于gin开启ws 去掉这三行则ws和http通过gin共用端口##############
	// http.HandleFunc("/ws", apis.WebsocketHandler)
	// go http.ListenAndServe(":6666", nil)
	// fmt.Println("Server is running: http://localhost:" + "6666")
	//#######################################################################
	//有订阅才开始抓第三方数据
	go func() {
		for {
			// dd := global.GD_WS.GetSubsTopics()
			// fmt.Println(dd)
			if len(global.GD_WS.Subscriptions) > 0 {
				// if global.GD_WS.GetSubsTopics().Contains(wsserver.PENDING) {
				// }
				// if global.GD_WS.GetSubsTopics().Contains(wsserver.GAS) {
				// }
				// if global.GD_WS.GetSubsTopics().Contains(wsserver.CoinPrice) {
				// }
				go jobs.FetchGas()
				go jobs.FetchPrice()
				// go jobs.FetchPending()
				//开始派发
				go global.GD_WS.Distribute()
				break
			}
			time.Sleep(time.Second * 1)
		}
	}()
	port := fmt.Sprintf("%d", global.GD_CONFIG.System.Port)
	fmt.Println(utils.Green("Websocket run at:"))
	fmt.Printf("-  Local:   http://localhost:%s/system/ws \r\n", port)
	fmt.Printf("-  Network: http://%s:%s/system/ws \r\n", utils.GetLocaHonst(), port)
	//后期应该用新线程监测是否有订阅，没有的时候就关闭抓取
}

func initChainProvider() {
	httpConn, err := ethclient.Dial(global.GD_CONFIG.ApiKeys.ProviderUrlHttps)
	if err == nil {
		fmt.Println(utils.Green("\r\n以太坊节点服务-HTTPS-连接正常"))
	}

	rpcCli, err2 := rpc.Dial(global.GD_CONFIG.ApiKeys.ProviderUrlHttps)
	if err2 == nil {
		fmt.Println(utils.Green("\r\n以太坊节点服务-WSS-连接正常"))
	}
	gcli := gethclient.New(rpcCli)
	// fmt.Println(global.GD_PROVIDERS.EtherscanClient.AccountBalance("0x1c4E5c4c6F8AaDF5F3B90EcE84F750f7293c188a"))
	global.GD_PROVIDERS.ProviderHttps = httpConn
	global.GD_PROVIDERS.ProviderWss = gcli
	global.GD_PROVIDERS.EtherscanClient = etherscan.New(etherscan.Mainnet, global.GD_CONFIG.ApiKeys.EtherScanKey)
	if nil != global.GD_PROVIDERS.EtherscanClient {
		fmt.Println(utils.Green("\r\nEtherScan-连接正常"))
	}
}

func otherInit() {
	dr, err := utils.ParseDuration(global.GD_CONFIG.JWT.ExpiresTime)
	if err != nil {
		panic(err)
	}
	_, err = utils.ParseDuration(global.GD_CONFIG.JWT.BufferTime)
	if err != nil {
		panic(err)
	}

	global.BlackCache = local_cache.NewCache(
		local_cache.SetDefaultExpire(dr),
	)
}
