package system

import (
	"context"
	. "golddigger/app/system/models"
	system "golddigger/cmd/initdb"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const initOrderMenu = initOrderAuthority + 1

type initMenu struct{}

// auto run
func init() {
	system.RegisterInit(initOrderMenu, &initMenu{})
}

func (i initMenu) InitializerName() string {
	return SysBaseMenu{}.TableName()
}

func (i *initMenu) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(
		&SysBaseMenu{},
		&SysBaseMenuParameter{},
		// &SysBaseMenuBtn{},
	)
}

func (i *initMenu) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	m := db.Migrator()
	return m.HasTable(&SysBaseMenu{}) &&
		m.HasTable(&SysBaseMenuParameter{})
	// &&m.HasTable(&SysBaseMenuBtn{})
}

func (i *initMenu) InitializeData(ctx context.Context) (next context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	entities := []SysBaseMenu{

		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "liveMint", Name: "liveMint", Component: "views/nft/liveMint/liveMint.vue", Sort: 1, Meta: Meta{Title: "pending池", Level: 1, Icon: "例外", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "quickSell", Name: "quickSell", Component: "views/nft/quickSell/quickSell.vue", Sort: 2, Meta: Meta{Title: "一键挂单", Level: 1, Icon: "fa-solid fa-cart-flatbed", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "nft",Redirect:"hashMint", Name: "nft", Component: "views/nft/index.vue", Sort: 3, Meta: Meta{Title: "NFT", Level: 1, Icon: "fa-solid fa-file-image", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "3", Path: "freeMint", Name: "freeMint", Component: "views/nft/freeMint/freeMint.vue", Sort: 1, Meta: Meta{Title: "Free Mint", Level: 2, Icon: "fa-solid fa-circle", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "3", Path: "followMint", Name: "followMint", Component: "views/nft/followMint/followMint.vue", Sort: 2, Meta: Meta{Title: "Follow Mint", Level: 2, Icon: "fa-solid fa-circle", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "3", Path: "hashMint", Name: "hashMint", Component: "views/nft/hashMint/hashMint.vue", Sort: 3, Meta: Meta{Title: "Hash Mint", Level: 2, Icon: "fa-solid fa-circle", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "3", Path: "quickOffer", Name: "quickOffer", Component: "views/nft/quickOffer/quickOffer.vue", Sort: 4, Meta: Meta{Title: "一键Offer", Level: 2, Icon: "fa-solid fa-circle", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "3", Path: "sweepFloor", Name: "sweepFloor", Component: "views/nft/sweepFloor/sweepFloor.vue", Sort: 5, Meta: Meta{Title: "扫地板", Level: 2, Icon: "fa-solid fa-circle", KeepAlive: true}},

		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "erc20",Redirect:"hotToken", Name: "erc20", Component: "views/erc20/index.vue", Sort: 4, Meta: Meta{Title: "ERC20", Level: 1, Icon: "fa-solid fa-coins", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "9", Path: "hotToken", Name: "hotToken", Component: "views/erc20/hotToken/hotToken.vue", Sort: 1, Meta: Meta{Title: "热门T代币", Level: 2, Icon: "fa-solid fa-circle", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "9", Path: "bigFlow", Name: "bigFlow", Component: "views/erc20/bigFlow/bigFlow.vue", Sort: 2, Meta: Meta{Title: "大额监控", Level: 2, Icon: "fa-solid fa-circle", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "9", Path: "tokenAnalysis", Name: "tokenAnalysis", Component: "views/erc20/tokenAnalysis/tokenAnalysis.vue", Sort: 3, Meta: Meta{Title: "代币分析", Level: 2, Icon: "fa-solid fa-circle", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "9", Path: "walletAnalysis", Name: "walletAnalysis", Component: "views/erc20/walletAnalysis/walletAnalysis.vue", Sort: 4, Meta: Meta{Title: "钱包分析", Level: 2, Icon: "fa-solid fa-circle", KeepAlive: true}},

		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "commonTools",Redirect:"tokenManage", Name: "commonTools", Component: "views/commonTools/commonTools.vue", Sort: 5, Meta: Meta{Title: "常用工具", Level: 1, Icon: "fa-solid fa-toolbox", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "14", Path: "frontRun", Name: "frontRun", Component: "views/commonTools/FrontRun/FrontRun.vue", Sort: 1, Meta: Meta{Title: "抢跑机器人", Level: 2, Icon: "fa-solid fa-circle", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "14", Path: "tokenManage", Name: "tokenManage", Component: "views/commonTools/tokenManage/tokenManage.vue", Sort: 2, Meta: Meta{Title: "代币分发/归集", Level: 2, Icon: "fa-solid fa-circle", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "14", Path: "nftManage", Name: "nftManage", Component: "views/commonTools/nftManage/nftManage.vue", Sort: 2, Meta: Meta{Title: "NFT分发/归集", Level: 2, Icon: "fa-solid fa-circle", KeepAlive: true}},

		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "admin",Redirect:"user", Name: "superAdmin", Component: "views/superAdmin/index.vue", Sort: 6, Meta: Meta{Title: "超级管理员", Level: 1, Icon: "fa-solid fa-user-secret"}},
		{MenuLevel: 0, Hidden: false, ParentId: "18", Path: "authority", Name: "authority", Component: "views/superAdmin/authority/authority.vue", Sort: 1, Meta: Meta{Title: "角色管理", Level: 2, Icon: "fa-solid fa-person-circle-question"}},
		{MenuLevel: 0, Hidden: false, ParentId: "18", Path: "menu", Name: "menu", Component: "views/superAdmin/menu/menu.vue", Sort: 2, Meta: Meta{Title: "菜单管理", Level: 2, Icon: "fa-solid fa-bars", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "18", Path: "api", Name: "api", Component: "views/superAdmin/api/api.vue", Sort: 3, Meta: Meta{Title: "api管理", Level: 2, Icon: "fa-solid fa-sliders", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "18", Path: "user", Name: "user", Component: "views/superAdmin/user/user.vue", Sort: 4, Meta: Meta{Title: "用户管理", Level: 2, Icon: "fa-solid fa-users"}},

		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "userFeedBack", Name: "userFeedBack", Component: "views/feedBack/userFeedBack.vue", Sort: 7, Meta: Meta{Title: "快速反馈", Level: 1, Icon: "fa-feather-pointed", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "adminFeedBack", Name: "adminFeedBack", Component: "views/feedBack/adminFeedBack.vue", Sort: 8, Meta: Meta{Title: "快速反馈", Icon: "fa-feather-pointed", KeepAlive: false}},

		{MenuLevel: 0, Hidden: true, ParentId: "0", Path: "upgrade", Name: "upgrade", Component: "views/account/upgrade/upgrade.vue", Sort: 8, Meta: Meta{Title: "账户升级", Icon: "fa-feather-pointed", KeepAlive: false}},
	}
	if err = db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, SysBaseMenu{}.TableName()+"表数据初始化失败!")
	}
	next = context.WithValue(ctx, i.InitializerName(), entities)
	return next, nil
}

func (i *initMenu) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.Where("path = ?", "autoPkg").First(&SysBaseMenu{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
