package router

import (
	v1 "golddigger/app/system/apis"
	"golddigger/middleware"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func /* (s *UserRouter) */ initUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user").Use(middleware.OperationRecord())
	userRouterWithoutRecord := Router.Group("user")
	userApi := v1.SystemApis.UserApi
	{
		userRouter.POST("setProviderUrl", userApi.SetProviderUrl)
		userRouter.POST("setTheme", userApi.SetTheme)
		userRouter.POST("setProviderDefault", userApi.SetProviderDefault)
		userRouter.POST("upgrade", userApi.Upgrade)
		// userRouter.POST("getNonce", baseApi.GetRandomSignMsg) // 管理员注册账号
		// userRouter.POST("admin_register", baseApi.Register) // 管理员注册账号
		// userRouter.POST("changePassword", baseApi.ChangePassword)         // 用户修改密码
		// userRouter.POST("setUserAuthority", baseApi.SetUserAuthority)     // 设置用户权限
		// userRouter.DELETE("deleteUser", baseApi.DeleteUser)               // 删除用户
		// userRouter.PUT("setUserInfo", baseApi.SetUserInfo)                // 设置用户信息
		// userRouter.PUT("setSelfInfo", baseApi.SetSelfInfo)                // 设置自身信息
		// userRouter.POST("setUserAuthorities", baseApi.SetUserAuthorities) // 设置用户权限组
		// userRouter.POST("resetPassword", baseApi.ResetPassword)           // 设置用户权限组
	}
	{
		// userRouterWithoutRecord.POST("getUserList", baseApi.GetUserList) // 分页获取用户列表
		userRouterWithoutRecord.GET("getUserInfo", userApi.GetUserInfo)  // 获取自身信息
		// userRouterWithoutRecord.POST("getNonce", baseApi.GetRandomSignMsg)
		// userRouterWithoutRecord.POST("login", baseApi.SignAndLogin)
	}
}
