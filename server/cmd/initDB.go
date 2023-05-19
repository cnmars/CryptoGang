/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"golddigger/app/system/models"
	"golddigger/cmd/initdb"
	"golddigger/global"
	ini "golddigger/init"
	"golddigger/utils"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
	_"golddigger/cmd/initdb/data"//导入只为执行包里面的init加载数据方法
)

// initDBCmd represents the initDB command
var (
	dbName    string
	userName  string
	passwd    string
	host      string
	port      string
	initDBCmd = &cobra.Command{
		Use:   "initdb",
		Short: "init database",
		Long:  `initdb是在一个新的环境还没有创建数据库的情况下使用,假如配置文件中的数据库能连通，则会提示数据库已存在（注意：用户名密码错误的情况下不能连通会视为数据库不存在）`,
		Example :"golddigger initdb -i 127.0.0.1 -o 3306 -d golddigger -u root -p 123456",
		Run: func(cmd *cobra.Command, args []string) {
			if dbName == "" {
				fmt.Println("没有指定数据库名")
				return
			}
			if userName == "" {
				fmt.Println("没有指定用户名")
				return
			}
			if passwd == "" {
				fmt.Println("没有指定密码")
				return
			}
			if host == "" {
				fmt.Println("没有指定主机位置")
				return
			}
			if port == "" {
				fmt.Println("没有指定数据库端口")
				return
			}
			initDB()
		},
	}
)

func init() {
	// StartCmd.Flags().StringVarP(&configYml, "config1", "c", "config/config.yaml", "Start server with provided configuration file")
	initDBCmd.Flags().StringVarP(&host, "host", "i", "", "password")
	initDBCmd.Flags().StringVarP(&port, "port", "o", "", "password")
	initDBCmd.Flags().StringVarP(&dbName, "dbname", "d", "", "database name")
	initDBCmd.Flags().StringVarP(&userName, "username", "u", "", "database username")
	initDBCmd.Flags().StringVarP(&passwd, "passwd", "p", "", "password")
	initDBCmd.MarkFlagRequired(dbName)
	initDBCmd.MarkFlagRequired(userName)
	initDBCmd.MarkFlagRequired(passwd)
	initDBCmd.MarkFlagRequired(host)
	initDBCmd.MarkFlagRequired(port)
}

func initDB() {
	ini.InitSystemForCMD(configYml)
	inUse := checkPort(global.GD_CONFIG.System.Port)
	if inUse {
		//线上环境在运行时不能初始化或者迁移数据库
		fmt.Println(utils.Red("服务器运行期间不能初始化数据库"))
	} else {
		if global.GD_DB != nil {
			global.GD_LOG.Error("已存在数据库配置!(配置文件中的数据库名已存在)")
			fmt.Println(utils.Red("已存在数据库配置！(配置文件中的数据库名已存在)"))
			return
		}
		var dbInfo = models.InitDB{
			DBType:   global.GD_CONFIG.System.DbType,
			Host:     host,
			Port:     port,
			UserName: userName,
			Password: passwd,
			DBName:   dbName,
		}
		if err := initdb.InitDB(dbInfo); err != nil {
			global.GD_LOG.Error("自动创建数据库失败!", zap.Error(err))
			fmt.Println(utils.Red("自动创建数据库失败，请查看后台日志，检查后在进行初始化"))
			return
		}
		fmt.Println(utils.Green("自动创建数据库成功"))
	}
}
