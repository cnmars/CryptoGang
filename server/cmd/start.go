package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	// "golddigger/common/wsserver"
	"golddigger/global"
	ini "golddigger/init"
	"golddigger/utils"

	// internal "golddigger/init/interobj"

	"github.com/spf13/cobra"
	// "go.uber.org/zap"
)

var (
	configYml string
	// apiCheck  bool
	StartCmd = &cobra.Command{
		Use:          "start",
		Short:        "Start API server",
		Example:      "golddigger start -c config/config.yaml",
		SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

var AppRouters = make([]func(), 0)

func init() {
	// StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config/config.yaml", "Start server with provided configuration file")
}

func setup() {

	ini.InitSystemOnline(configYml)
	// usageStr := `starting api server...`
	// ini.RunWindowsServer()
	// log.Println(usageStr)
}

func run() error {
	// go func() {
	// 	jobs.InitJob()
	// 	jobs.Setup(sdk.Runtime.GetDb())

	// }()
	Router := global.GD_ENGIN
	//静态文件路由
	// Router.Static("/form-generator", "./resource/page")
	address := fmt.Sprintf(":%d", global.GD_CONFIG.System.Port)
	// srv := internal.InitServer(address, Router)
	srv := &http.Server{
		Addr:    address,
		Handler: Router,
		// ReadTimeout:    20 * time.Second,
		// WriteTimeout:   20 * time.Second,
		// MaxHeaderBytes: 1 << 20,
	}

	go func() {
		// 服务连接
		if global.GD_CONFIG.SslConfig.Enable {
			if err := srv.ListenAndServeTLS(global.GD_CONFIG.SslConfig.Pem, global.GD_CONFIG.SslConfig.KeyStr); err != nil && err != http.ErrServerClosed {
				log.Fatal("listen: ", err)
			}
		} else {
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatal("listen: ", err)
			}
		}
	}()
	// fmt.Println(pkg.Red(string(global.LogoContent)))

	// global.GD_LOG.Info("server run at: ", zap.String("address", port))
	fmt.Println(utils.Green("Server run at:"))
	fmt.Printf("-  Local:   http://localhost%s/ \r\n", address)
	fmt.Printf("-  Network: http://%s%s/ \r\n", utils.GetLocaHonst(), address)
	// fmt.Println("Swagger run at:")
	// fmt.Printf("-  Local:   http://localhost:%d/swagger/admin/index.html \r\n", config.ApplicationConfig.Port)
	// fmt.Printf("-  Network: http://%s:%d/swagger/admin/index.html \r\n", pkg.GetLocaHonst(), config.ApplicationConfig.Port)
	tip()
	//启动websocket
	// wsserver.StartWsService()
	fmt.Printf("Enter Control + C Shutdown Server \r\n")
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	fmt.Printf("Shutdown Server ... \r\n")

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
	return nil
}

func tip() {
	fmt.Println()
	usageStr := `=========欢迎使用 ` + utils.Green(`GoldGigger `+global.Version) + ` 可以使用 ` + utils.Green(`-h`) + ` 查看命令=========`
	fmt.Printf("%s \n\n", usageStr)
}
