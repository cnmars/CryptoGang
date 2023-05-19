package system

import (
	"context"
	sysModel "golddigger/app/system/models"
	system"golddigger/cmd/initdb"
	"golddigger/utils"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const initOrderAuthority = initOrderCasbin + 1

type initAuthority struct{}

// auto run
func init() {
	system.RegisterInit(initOrderAuthority, &initAuthority{})
}

func (i *initAuthority) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&sysModel.SysAuthority{})
}

func (i *initAuthority) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&sysModel.SysAuthority{})
}

func (i initAuthority) InitializerName() string {
	return sysModel.SysAuthority{}.TableName()
}

func (i *initAuthority) InitializeData(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	entities := []sysModel.SysAuthority{
		{AuthorityId: 888, AuthorityName: "管理员", ParentId: utils.Pointer[uint](0), DefaultRouter: "liveMint"},
		{AuthorityId: 333, AuthorityName: "普通用户", ParentId: utils.Pointer[uint](0), DefaultRouter: "liveMint"},
		{AuthorityId: 666, AuthorityName: "付费用户", ParentId: utils.Pointer[uint](0), DefaultRouter: "liveMint"},
	}

	if err := db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrapf(err, "%s表数据初始化失败!", sysModel.SysAuthority{}.TableName())
	}
	// data authority
	if err := db.Model(&entities[0]).Association("DataAuthorityId").Replace(
		[]*sysModel.SysAuthority{
			{AuthorityId: 888},
			{AuthorityId: 333},
			{AuthorityId: 666},
		}); err != nil {
		return ctx, errors.Wrapf(err, "%s表数据初始化失败!",
			db.Model(&entities[0]).Association("DataAuthorityId").Relationship.JoinTable.Name)
	}
	if err := db.Model(&entities[1]).Association("DataAuthorityId").Replace(
		[]*sysModel.SysAuthority{
			{AuthorityId: 333},
			{AuthorityId: 666},
		}); err != nil {
		return ctx, errors.Wrapf(err, "%s表数据初始化失败!",
			db.Model(&entities[1]).Association("DataAuthorityId").Relationship.JoinTable.Name)
	}

	next := context.WithValue(ctx, i.InitializerName(), entities)
	return next, nil
}

func (i *initAuthority) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.Where("authority_id = ?", "666").
		First(&sysModel.SysAuthority{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
