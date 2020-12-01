package server

import (
	"call-up/api"
	"call-up/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Cors())

	// JWT中间件
	authMiddleware, err := middleware.GinJWTMiddlewareInit()
	if err != nil {
		log.Fatal("JWT Error:", err.Error())
	}

	// When you use jwt.New(), the function is already automatically called for checking,
	// which means you don't need to call it again.
	if err := authMiddleware.MiddlewareInit(); err != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:", err.Error())
	}

	// 路由
	apis := r.Group("/api")
	{
		apis.POST("ping", api.Ping)

		// 用户注册、登录
		apis.POST("user/login", authMiddleware.LoginHandler)
		apis.POST("user/reg", api.UserRegister)

		// 需要登录保护的
		auth := apis.Group("")
		{
			auth.Use(authMiddleware.MiddlewareFunc())
			auth.GET("refresh", authMiddleware.RefreshHandler)
			auth.POST("user/logout", authMiddleware.LogoutHandler)
			auth.POST("user/info", api.UserInfo)
			auth.POST("user/updatePasswd", api.UserPasswordUpdate)
			auth.POST("user/updateInfo", api.UserInfoUpdate)
			// 需要普通用户权限的
			userAuth := auth.Group("")
			{
				userAuth.Use(middleware.UserAuth())
				userAuth.POST("callup/create", api.CallupCreate)
			}
			//需要管理员用户权限的
			adminAuth := auth.Group("")
			{
				adminAuth.Use(middleware.AdminAuth())
			}
		}

	}
	return r
}
