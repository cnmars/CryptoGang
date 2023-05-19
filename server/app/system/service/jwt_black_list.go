package service

import (
	"context"

	"go.uber.org/zap"

	"golddigger/app/system/models"
	system "golddigger/app/system/models"
	"golddigger/global"
	"golddigger/utils/jwtutil"
)

type JwtService struct{}

//@description: 拉黑jwt
//@param: jwtList model.JwtBlacklist
//@return: err error

func (jwtService *JwtService) JsonInBlacklist(jwtList models.JwtBlacklist) (err error) {
	err = global.GD_DB.Create(&jwtList).Error
	if err != nil {
		return
	}
	global.BlackCache.SetDefault(jwtList.Jwt, struct{}{})
	return
}

//@description: 判断JWT是否在黑名单内部
//@param: jwt string
//@return: bool

func (jwtService *JwtService) IsBlacklist(jwt string) bool {
	_, ok := global.BlackCache.Get(jwt)
	return ok
	// err := global.GD_DB.Where("jwt = ?", jwt).First(&system.JwtBlacklist{}).Error
	// isNotFound := errors.Is(err, gorm.ErrRecordNotFound)
	// return !isNotFound
}

//@description: 从redis取jwt
//@param: userName string

func (jwtService *JwtService) GetRedisJWT(userName string) (redisJWT string, err error) {
	redisJWT, err = global.GD_REDIS.Get(context.Background(), userName).Result()
	// redisJWT, err = global.GD_REDIS.Get( /* context.Background(), */ userName).Result()
	return redisJWT, err
}

//@description: jwt存入redis并设置过期时间
//@param: jwt string, userName string
//@return: err error

func (jwtService *JwtService) SetRedisJWT(jwt string, userName string) (err error) {
	// 此处过期时间等于jwt过期时间
	dr, err := jwtutil.ParseDuration(global.GD_CONFIG.JWT.ExpiresTime)
	if err != nil {
		return err
	}
	timer := dr
	// err = global.GD_REDIS.Set( /* context.Background(), */ userName, jwt, timer).Err()
	err = global.GD_REDIS.Set(context.Background(), userName, jwt, timer).Err()
	return err
}

func LoadAll() {
	var data []string
	err := global.GD_DB.Model(&system.JwtBlacklist{}).Select("jwt").Find(&data).Error
	if err != nil {
		global.GD_LOG.Error("加载数据库jwt黑名单失败!", zap.Error(err))
		return
	}
	for i := 0; i < len(data); i++ {
		global.BlackCache.SetDefault(data[i], struct{}{})
	} // jwt黑名单 加入 BlackCache 中
}
