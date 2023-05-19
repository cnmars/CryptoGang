/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"golddigger/global"
	"golddigger/utils"
	"net"
	"os"
	"strconv"

	adapter "github.com/casbin/gorm-adapter/v3"
	"go.uber.org/zap"

	// "gorm.io/gorm"
	. "golddigger/app/system/models"
	ini "golddigger/init"

	// user "golddigger/app/user/models"
	// common"golddigger/common/models"
	"github.com/spf13/cobra"
)

// 自动更新数据库表
var migrateCmd = &cobra.Command{
	Use:     "migrate",
	Short:   "自动迁移数据库",
	Long:    `migrate是在开发环境上数据库结构发生改变的时候使用`,
	Example: "golddigger migrate",
	Run: func(cmd *cobra.Command, args []string) {
		migrate()
	},
}

func init() {
	// StartCmd.PersistentFlags().StringVarP(&configYml, "config2", "c", "config/config.yaml", "Start server with provided configuration file")
}
func migrate() {

	ini.InitSystemForCMD(configYml)
	inUse := checkPort(global.GD_CONFIG.System.Port)
	if !inUse {
		if global.GD_DB != nil {
			fmt.Println(utils.Green("开始迁移"))
			registerTables() // 初始化表
			// insertData()     //初始化数据
			// 程序结束前关闭数据库链接
			db, _ := global.GD_DB.DB()
			defer db.Close()
		}

	} else {
		fmt.Println(utils.Red("系统运行期间不能迁移！"))
	}
}
func registerTables() {
	err := global.GD_DB.Debug().AutoMigrate(
		// 系统模块表
		SysUser{},
		JwtBlacklist{},
		SysAuthority{},
		SysOperationRecord{},
		FeedBack{},

		SysBaseMenu{},
		SysBaseMenuParameter{},
		SysAuthority{},
		adapter.CasbinRule{},
		UpgradeRecord{},
		// 示例模块表
	)
	if err != nil {
		global.GD_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.GD_LOG.Info("register table success")
}

func insertData() {
	//插入菜单数据
	entities := []SysBaseMenu{

		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "liveMint", Name: "liveMint", Component: "views/nft/liveMint/liveMint.vue", Sort: 1, Meta: Meta{Title: "pending池", Level: 1, Icon: "例外", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "quickSell", Name: "quickSell", Component: "views/nft/quickSell/quickSell.vue", Sort: 2, Meta: Meta{Title: "一键挂单", Level: 1, Icon: "fa-solid fa-cart-flatbed", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "nft", Name: "nft", Component: "views/nft/index.vue", Sort: 3, Meta: Meta{Title: "NFT", Level: 1, Icon: "fa-solid fa-file-image", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "3", Path: "freeMint", Name: "freeMint", Component: "views/nft/freeMint/freeMint.vue", Sort: 1, Meta: Meta{Title: "Free Mint", Level: 2, Icon: "fa-solid fa-circle", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "3", Path: "followMint", Name: "followMint", Component: "views/nft/followMint/followMint.vue", Sort: 2, Meta: Meta{Title: "Follow Mint", Level: 2, Icon: "fa-solid fa-circle", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "3", Path: "hashMint", Name: "hashMint", Component: "views/nft/hashMint/hashMint.vue", Sort: 3, Meta: Meta{Title: "Hash Mint", Level: 2, Icon: "fa-solid fa-circle", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "3", Path: "quickOffer", Name: "quickOffer", Component: "views/nft/quickOffer/quickOffer.vue", Sort: 4, Meta: Meta{Title: "一键Offer", Level: 2, Icon: "fa-solid fa-circle", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "3", Path: "sweepFloor", Name: "sweepFloor", Component: "views/nft/sweepFloor/sweepFloor.vue", Sort: 5, Meta: Meta{Title: "扫地板", Level: 2, Icon: "fa-solid fa-circle", KeepAlive: true}},

		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "erc20", Name: "ERC20", Component: "views/erc20/index.vue", Sort: 4, Meta: Meta{Title: "ERC20", Level: 1, Icon: "fa-solid fa-bitcoin", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "9", Path: "hotToken", Name: "hotToken", Component: "views/erc20/hotToken/hotToken.vue", Sort: 1, Meta: Meta{Title: "热门T代币", Level: 2, Icon: "fa-solid fa-bitcoin", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "9", Path: "bigFlow", Name: "bigFlow", Component: "views/erc20/bigFlow/bigFlow.vue", Sort: 2, Meta: Meta{Title: "大额监控", Level: 2, Icon: "fa-solid fa-bitcoin", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "9", Path: "tokenAnalysis", Name: "tokenAnalysis", Component: "views/erc20/tokenAnalysis/tokenAnalysis.vue", Sort: 3, Meta: Meta{Title: "代币分析", Level: 2, Icon: "fa-solid fa-bitcoin", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "9", Path: "walletAnalysis", Name: "walletAnalysis", Component: "views/erc20/walletAnalysis/walletAnalysis.vue", Sort: 4, Meta: Meta{Title: "钱包分析", Level: 2, Icon: "fa-solid fa-bitcoin", KeepAlive: true}},

		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "commonTool", Name: "commonTool", Component: "views/commonTool/commonTool.vue", Sort: 5, Meta: Meta{Title: "常用工具", Level: 1, Icon: "fa-solid fa-bitcoin", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "14", Path: "FrontRun", Name: "FrontRun", Component: "views/commonTool/FrontRun/FrontRun.vue", Sort: 1, Meta: Meta{Title: "抢跑机器人", Level: 2, Icon: "fa-solid fa-bitcoin", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "14", Path: "tokenManage", Name: "tokenManage", Component: "views/commonTool/tokenManage/tokenManage.vue", Sort: 2, Meta: Meta{Title: "代币分发/归集", Level: 2, Icon: "fa-solid fa-bitcoin", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "14", Path: "nftManage", Name: "nftManage", Component: "views/commonTool/nftManage/nftManage.vue", Sort: 2, Meta: Meta{Title: "NFT分发/归集", Level: 2, Icon: "fa-solid fa-bitcoin", KeepAlive: true}},

		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "userFeedBack", Name: "userFeedBack", Component: "views/feedBack/userFeedBack.vue", Sort: 6, Meta: Meta{Title: "快速反馈", Level: 1, Icon: "fa-feather-pointed", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "adminFeedBack", Name: "adminFeedBack", Component: "views/feedBack/adminFeedBack.vue", Sort: 6, Meta: Meta{Title: "快速反馈", Icon: "fa-feather-pointed", KeepAlive: false}},

		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "admin", Name: "superAdmin", Component: "views/superAdmin/index.vue", Sort: 7, Meta: Meta{Title: "超级管理员", Level: 1, Icon: "fa-solid fa-user-secret"}},
		{MenuLevel: 0, Hidden: false, ParentId: "20", Path: "authority", Name: "authority", Component: "views/superAdmin/authority/authority.vue", Sort: 1, Meta: Meta{Title: "角色管理", Level: 2, Icon: "fa-solid fa-person-circle-question"}},
		{MenuLevel: 0, Hidden: false, ParentId: "20", Path: "menu", Name: "menu", Component: "views/superAdmin/menu/menu.vue", Sort: 2, Meta: Meta{Title: "菜单管理", Level: 2, Icon: "fa-solid fa-bars", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "20", Path: "api", Name: "api", Component: "views/superAdmin/api/api.vue", Sort: 3, Meta: Meta{Title: "api管理", Level: 2, Icon: "fa-solid fa-sliders", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "20", Path: "user", Name: "user", Component: "views/superAdmin/user/user.vue", Sort: 4, Meta: Meta{Title: "用户管理", Level: 2, Icon: "fa-solid fa-users"}},
	}
	if err := global.GD_DB.Create(&entities).Error; err != nil {
		global.GD_LOG.Error("表数据初始化失败!", zap.Error(err))
		os.Exit(0)
	}
	//插入菜单权限数据
	// 888
	if err := global.GD_DB.Model(&SysAuthority{}).Where("authority_id = ?", 888).Association("SysBaseMenus").Append(entities[:24]); err != nil {
		global.GD_LOG.Error("插入888菜单权限数据失败!", zap.Error(err))
	}
	// 666
	if err := global.GD_DB.Model(&SysAuthority{}).Where("authority_id = ?", 666).Association("SysBaseMenus").Append(entities[:18]); err != nil {
		global.GD_LOG.Error("插入666菜单权限数据失败!", zap.Error(err))
	}
	// 333
	if err := global.GD_DB.Model(&SysAuthority{}).Where("authority_id = ?", 333).Association("SysBaseMenus").Append(entities[:12]); err != nil {
		global.GD_LOG.Error("插入333菜单权限数据失败!", zap.Error(err))
	}
	global.GD_LOG.Info("表数据初始化完成!")
}

func checkPort(port int) (status bool) {

	// Concatenate a colon and the port
	host := ":" + strconv.Itoa(port)

	// Try to create a server with the port
	server, err := net.Listen("tcp", host)

	// if it fails then the port is likely taken
	if err != nil {
		return true
	}

	// close the server
	server.Close()

	// we successfully used and closed the port
	// so it's now available to be used again
	return false

}
