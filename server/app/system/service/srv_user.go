package service

import (
	"errors"
	system "golddigger/app/system/models"
	"golddigger/app/system/models/request/common"
	"golddigger/global"
	"golddigger/utils"
	"time"

	// uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserService struct{}

func (userService *UserService) Login(u *system.SysUser) (userInter *system.SysUser, ty int) {
	if nil == global.GD_DB {
		global.GD_LOG.Error("数据库未初始化！")
		return nil, 0
	}

	var user system.SysUser
	err := global.GD_DB.Where("address = ?", u.Address).Preload("Authorities").Preload("Authority").First(&user).Error
	if err != nil {
		global.GD_LOG.Info("新用户："+u.Address+" 注册", zap.Error(err))
		user.RecentLogin = time.Now()
		user.Address = u.Address
		err = global.GD_DB.Create(&user).Error
		if err != nil {
			global.GD_LOG.Error("新用户："+u.Address+" 注册失败", zap.Error(err))
			return &user, 1
		}
	}
	return &user, 2
}

func (userService *UserService) GetUserInfoList(info common.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GD_DB.Model(&system.SysUser{})
	var userList []system.SysUser
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Preload("Authorities").Preload("Authority").Find(&userList).Error
	return userList, total, err
}

func (userService *UserService) SetUserAuthority(id uint, authorityId uint) (err error) {
	assignErr := global.GD_DB.Where("sys_user_id = ? AND sys_authority_authority_id = ?", id, authorityId).First(&system.SysUserAuthority{}).Error
	if errors.Is(assignErr, gorm.ErrRecordNotFound) {
		return errors.New("该用户无此角色")
	}
	err = global.GD_DB.Where("id = ?", id).First(&system.SysUser{}).Update("authority_id", authorityId).Error
	return err
}

func (userService *UserService) SetUserAuthorities(id uint, authorityIds []uint) (err error) {
	return global.GD_DB.Transaction(func(tx *gorm.DB) error {
		TxErr := tx.Delete(&[]system.SysUserAuthority{}, "sys_user_id = ?", id).Error
		if TxErr != nil {
			return TxErr
		}
		var useAuthority []system.SysUserAuthority
		for _, v := range authorityIds {
			useAuthority = append(useAuthority, system.SysUserAuthority{
				SysUserId: id, SysAuthorityAuthorityId: v,
			})
		}
		TxErr = tx.Create(&useAuthority).Error
		if TxErr != nil {
			return TxErr
		}
		TxErr = tx.Where("id = ?", id).First(&system.SysUser{}).Update("authority_id", authorityIds[0]).Error
		if TxErr != nil {
			return TxErr
		}
		// 返回 nil 提交事务
		return nil
	})
}

func (userService *UserService) DeleteUser(id int) (err error) {
	var user system.SysUser
	err = global.GD_DB.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return err
	}
	err = global.GD_DB.Delete(&[]system.SysUserAuthority{}, "sys_user_id = ?", id).Error
	return err
}

func (userService *UserService) SetUserInfo(req system.SysUser) error {
	return global.GD_DB.Updates(&req).Error
}

func (userService *UserService) GetUserInfo(address string) (user system.SysUser, err error) {
	var reqUser system.SysUser
	err = global.GD_DB.Preload("Authorities").Preload("Authority").First(&reqUser, "address = ?", address).Error
	if err != nil {
		return reqUser, err
	}
	if reqUser.AuthorityId == 666 {
		now := time.Now()
		if now.After(reqUser.ExpireAt) {
			err = global.GD_DB.Model(&system.SysUser{}).Where("address = ?", address).Updates(system.SysUser{AuthorityId: 333}).Error
			reqUser.AuthorityId = 333
			return reqUser, err
		}
	}
	// MenuServiceApp.UserAuthorityDefaultRouter(&reqUser)
	return reqUser, err
}

func (userService *UserService) SetHttpProvider(address string, url string) (err error) {
	data := make(map[string]interface{})
	data["http_provider"] = url
	data["use_default_http"] = 0 //零值字段
	err = global.GD_DB.Where("address = ?", address).First(&system.SysUser{}).Updates(data).Error
	return err
}

func (userService *UserService) SetHttpDefault(address string) (err error) {
	err = global.GD_DB.Where("address = ?", address).First(&system.SysUser{}).Updates(system.SysUser{UseDefaultHttp: true}).Error
	return err
}
func (userService *UserService) SetTheme(address string,theme string) (err error) {
	err = global.GD_DB.Where("address = ?", address).First(&system.SysUser{}).Updates(system.SysUser{SiteTheme: theme}).Error
	return err
}

func (userService *UserService) SetWssProvider(address string, url string) (err error) {
	data := make(map[string]interface{})
	data["wss_provider"] = url
	data["use_default_wss"] = 0 //零值字段
	err = global.GD_DB.Where("address = ?", address).First(&system.SysUser{}).Updates(data).Error
	return err
}

func (userService *UserService) SetWssDefault(address string) (err error) {
	err = global.GD_DB.Where("address = ?", address).First(&system.SysUser{}).Updates(system.SysUser{UseDefaultWss: true}).Error
	return err
}

func (userService *UserService) FindUserById(id int) (user *system.SysUser, err error) {
	var u system.SysUser
	err = global.GD_DB.Where("`id` = ?", id).First(&u).Error
	return &u, err
}

func (userService *UserService) FindUserByUuid(uuid string) (user *system.SysUser, err error) {
	var u system.SysUser
	if err = global.GD_DB.Where("`uuid` = ?", uuid).First(&u).Error; err != nil {
		return &u, errors.New("用户不存在")
	}
	return &u, nil
}

func (userService *UserService) ResetPassword(ID uint) (err error) {
	err = global.GD_DB.Model(&system.SysUser{}).Where("id = ?", ID).Update("password", utils.BcryptHash("123456")).Error
	return err
}

func (userService *UserService) Upgrade(upgradeType uint16, hash string, address string) (user system.SysUser, err error) {
	var now = time.Now()
	var u system.SysUser
	err = global.GD_DB.Where("address = ?", address).First(&u).Error
	if err != nil {
		return u, err
	}
	var expireAt time.Time
	switch upgradeType {
	case 1:
		expireAt = now.AddDate(0, 1, 0)
	case 2:
		expireAt = now.AddDate(1, 0, 0)
	}
	var record system.UpgradeRecord
	record.Type = upgradeType
	record.TranscationHash = hash
	record.Address = u.Address
	record.ExpireAt = expireAt
	err = global.GD_DB.Model(&system.SysUser{}).Where("address = ?", address).Updates(system.SysUser{ExpireAt: expireAt, AuthorityId: 666}).Error
	if err != nil {
		return u, err
	}
	err = global.GD_DB.Create(&record).Error
	u.AuthorityId = 666
	return u, err
}
