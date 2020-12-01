package server

import (
	"call-up/api"
	"call-up/middleware"
	"github.com/gin-gonic/gin"
	"log"
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


		// 需要用户登录保护的
		userAuth := apis.Group("")
		{
			userAuth.Use(authMiddleware.MiddlewareFunc())
			userAuth.GET("refresh", authMiddleware.RefreshHandler)
			userAuth.POST("user/logout", authMiddleware.LogoutHandler)
			userAuth.POST("user/info", api.UserInfo)
			userAuth.POST("user/updatePasswd", api.UserPasswordUpdate)
			userAuth.POST("user/updateInfo", api.UserInfoUpdate)
		}
	}
	return r
}
