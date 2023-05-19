package service

import (
	"errors"
	"strconv"

	system "golddigger/app/system/models"
	"golddigger/app/system/models/request/common"
	"golddigger/app/system/models/response"
	"golddigger/global"

	"gorm.io/gorm"
)

var ErrRoleExistence = errors.New("存在相同角色id")

//@description: 创建一个角色
//@param: auth model.SysAuthority
//@return: authority system.SysAuthority, err error

type AuthorityService struct{}

var AuthorityServiceApp = new(AuthorityService)

func (authorityService *AuthorityService) CreateAuthority(auth system.SysAuthority) (authority system.SysAuthority, err error) {
	var authorityBox system.SysAuthority
	if !errors.Is(global.GD_DB.Where("authority_id = ?", auth.AuthorityId).First(&authorityBox).Error, gorm.ErrRecordNotFound) {
		return auth, ErrRoleExistence
	}
	err = global.GD_DB.Create(&auth).Error
	return auth, err
}

//@description: 复制一个角色
//@param: copyInfo response.SysAuthorityCopyResponse
//@return: authority system.SysAuthority, err error

func (authorityService *AuthorityService) CopyAuthority(copyInfo response.SysAuthorityCopyResponse) (authority system.SysAuthority, err error) {
	var authorityBox system.SysAuthority
	if !errors.Is(global.GD_DB.Where("authority_id = ?", copyInfo.Authority.AuthorityId).First(&authorityBox).Error, gorm.ErrRecordNotFound) {
		return authority, ErrRoleExistence
	}
	copyInfo.Authority.Children = []system.SysAuthority{}
	if err != nil {
		return
	}
	err = global.GD_DB.Create(&copyInfo.Authority).Error
	if err != nil {
		return
	}


	if err != nil {
		return
	}
	paths := CasbinServiceApp.GetPolicyPathByAuthorityId(copyInfo.OldAuthorityId)
	err = CasbinServiceApp.UpdateCasbin(copyInfo.Authority.AuthorityId, paths)
	if err != nil {
		_ = authorityService.DeleteAuthority(&copyInfo.Authority)
	}
	return copyInfo.Authority, err
}

//@description: 更改一个角色
//@param: auth model.SysAuthority
//@return: authority system.SysAuthority, err error

func (authorityService *AuthorityService) UpdateAuthority(auth system.SysAuthority) (authority system.SysAuthority, err error) {
	err = global.GD_DB.Where("authority_id = ?", auth.AuthorityId).First(&system.SysAuthority{}).Updates(&auth).Error
	return auth, err
}

//@description: 删除角色
//@param: auth *model.SysAuthority
//@return: err error

func (authorityService *AuthorityService) DeleteAuthority(auth *system.SysAuthority) (err error) {
	if errors.Is(global.GD_DB.Debug().Preload("Users").First(&auth).Error, gorm.ErrRecordNotFound) {
		return errors.New("该角色不存在")
	}
	if len(auth.Users) != 0 {
		return errors.New("此角色有用户正在使用禁止删除")
	}
	if !errors.Is(global.GD_DB.Where("authority_id = ?", auth.AuthorityId).First(&system.SysUser{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("此角色有用户正在使用禁止删除")
	}
	if !errors.Is(global.GD_DB.Where("parent_id = ?", auth.AuthorityId).First(&system.SysAuthority{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("此角色存在子角色不允许删除")
	}
	db := global.GD_DB.Preload("SysBaseMenus").Where("authority_id = ?", auth.AuthorityId).First(auth)
	err = db.Unscoped().Delete(auth).Error
	if err != nil {
		return
	}
	err = global.GD_DB.Delete(&[]system.SysUserAuthority{}, "sys_authority_authority_id = ?", auth.AuthorityId).Error
	authorityId := strconv.Itoa(int(auth.AuthorityId))
	CasbinServiceApp.ClearCasbin(0, authorityId)
	return err
}

//@description: 分页获取数据
//@param: info common.PageInfo
//@return: list interface{}, total int64, err error

func (authorityService *AuthorityService) GetAuthorityInfoList(info common.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GD_DB.Model(&system.SysAuthority{})
	err = db.Where("parent_id = ?", "0").Count(&total).Error
	var authority []system.SysAuthority
	err = db.Limit(limit).Offset(offset).Preload("DataAuthorityId").Where("parent_id = ?", "0").Find(&authority).Error
	if len(authority) > 0 {
		for k := range authority {
			err = authorityService.findChildrenAuthority(&authority[k])
		}
	}
	return authority, total, err
}

//@description: 获取所有角色信息
//@param: auth model.SysAuthority
//@return: sa system.SysAuthority, err error

func (authorityService *AuthorityService) GetAuthorityInfo(auth system.SysAuthority) (sa system.SysAuthority, err error) {
	err = global.GD_DB.Preload("DataAuthorityId").Where("authority_id = ?", auth.AuthorityId).First(&sa).Error
	return sa, err
}

//@description: 设置角色资源权限
//@param: auth model.SysAuthority
//@return: error

func (authorityService *AuthorityService) SetDataAuthority(auth system.SysAuthority) error {
	var s system.SysAuthority
	global.GD_DB.Preload("DataAuthorityId").First(&s, "authority_id = ?", auth.AuthorityId)
	err := global.GD_DB.Model(&s).Association("DataAuthorityId").Replace(&auth.DataAuthorityId)
	return err
}

//@description: 菜单与角色绑定
//@param: auth *model.SysAuthority
//@return: error
func (authorityService *AuthorityService) SetMenuAuthority(auth *system.SysAuthority) error {
	var s system.SysAuthority
	global.GD_DB.Preload("SysBaseMenus").First(&s, "authority_id = ?", auth.AuthorityId)
	err := global.GD_DB.Model(&s).Association("SysBaseMenus").Replace(&auth.SysBaseMenus)
	return err
}
//@description: 查询子角色
//@param: authority *model.SysAuthority
//@return: err error

func (authorityService *AuthorityService) findChildrenAuthority(authority *system.SysAuthority) (err error) {
	err = global.GD_DB.Preload("DataAuthorityId").Where("parent_id = ?", authority.AuthorityId).Find(&authority.Children).Error
	if len(authority.Children) > 0 {
		for k := range authority.Children {
			err = authorityService.findChildrenAuthority(&authority.Children[k])
		}
	}
	return err
}
