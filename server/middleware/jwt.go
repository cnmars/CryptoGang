package middleware

import (
	"strconv"
	"time"

	"golddigger/utils/jwtutil"

	system "golddigger/app/system/models"
	"golddigger/app/system/service"
	response "golddigger/common/models/response"
	"golddigger/global"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var jwtService = service.SystemService.JwtService

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := c.Request.Header.Get("x-token")
		if token == "" {
			//websocket的token
			token = c.Request.Header.Get("Sec-WebSocket-Protocol")
			if token == "" {
				response.FailWithDetailed(gin.H{"reload": true}, "请使用钱包登录！", c)
				c.Abort()
				return
			}
		}
		if jwtService.IsBlacklist(token) {
			response.FailWithDetailed(gin.H{"reload": true}, "您的帐户异地登陆或令牌失效", c)
			c.Abort()
			return
		}
		j := jwtutil.NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == jwtutil.TokenExpired {
				response.FailWithDetailed(gin.H{"reload": true}, "授权已过期", c)
				c.Abort()
				return
			}
			response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
			c.Abort()
			return
		}

		// 已登录用户被管理员禁用 需要使该用户的jwt失效 此处比较消耗性能 如果需要 请自行打开
		// 用户被删除的逻辑 需要优化 此处比较消耗性能 如果需要 请自行打开

		//if user, err := userService.FindUserByUuid(claims.UUID.String()); err != nil || user.Enable == 2 {
		//	_ = jwtService.JsonInBlacklist(system.JwtBlacklist{Jwt: token})
		//	response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
		//	c.Abort()
		//}
		if claims.ExpiresAt-time.Now().Unix() < claims.BufferTime {
			dr, _ := jwtutil.ParseDuration(global.GD_CONFIG.JWT.ExpiresTime)
			claims.ExpiresAt = time.Now().Add(dr).Unix()
			newToken, _ := j.CreateTokenByOldToken(token, *claims)
			global.GD_LOG.Info("token已过期，新token:")
			global.GD_LOG.Info(newToken)
			newClaims, _ := j.ParseToken(newToken)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt, 10))
			if global.GD_CONFIG.System.UseMultipoint {
				RedisJwtToken, err := jwtService.GetRedisJWT(newClaims.Address)
				if err != nil {
					global.GD_LOG.Error("get redis jwt failed", zap.Error(err))
				} else { // 当之前的取成功时才进行拉黑操作
					_ = jwtService.JsonInBlacklist(system.JwtBlacklist{Jwt: RedisJwtToken})
				}
				// 无论如何都要记录当前的活跃状态
				_ = jwtService.SetRedisJWT(newToken, newClaims.Address)
			}
		}
		c.Set("claims", claims)
		c.Next()
	}
}
