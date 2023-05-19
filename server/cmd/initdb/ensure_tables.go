package initdb

import (
	"context"
	adapter "github.com/casbin/gorm-adapter/v3"
	system "golddigger/app/system/models"
	// user "golddigger/app/user/models"
	// common "golddigger/common/models"
	"gorm.io/gorm"
)

const initOrderEnsureTables = InitOrderExternal - 1

type ensureTables struct{}

// auto run
func init() {
	RegisterInit(initOrderEnsureTables, &ensureTables{})
}

func (ensureTables) InitializerName() string {
	return "ensure_tables_created"
}
func (e *ensureTables) InitializeData(ctx context.Context) (next context.Context, err error) {
	return ctx, nil
}

func (e *ensureTables) DataInserted(ctx context.Context) bool {
	return true
}

func (e *ensureTables) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, ErrMissingDBContext
	}
	tables := []interface{}{
		system.SysUser{},
		system.SysBaseMenu{},
		system.SysBaseMenuParameter{},
		system.SysAuthority{},

		system.JwtBlacklist{},
		system.SysOperationRecord{},
		system.FeedBack{},
		adapter.CasbinRule{},
		system.UpgradeRecord{},

	}
	for _, t := range tables {
		_ = db.AutoMigrate(&t)
		// 视图 authority_menu 会被当成表来创建，引发冲突错误（更新版本的gorm似乎不会）
		// 由于 AutoMigrate() 基本无需考虑错误，因此显式忽略
	}
	return ctx, nil
}

func (e *ensureTables) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	tables := []interface{}{
		system.SysUser{},
		system.SysBaseMenu{},
		system.SysBaseMenuParameter{},
		system.SysAuthority{},

		system.JwtBlacklist{},
		system.SysOperationRecord{},
		system.FeedBack{},
		adapter.CasbinRule{},
		system.UpgradeRecord{},

	}
	yes := true
	for _, t := range tables {
		yes = yes && db.Migrator().HasTable(t)
	}
	return yes
}
