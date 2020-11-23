package server

import (
	"call-up/api"
	"call-up/middleware"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())

	// JWT中间件
	authMiddleware, err := middleware.GinJWTMiddlewareInit()
	if err != nil {
		log.Fatal("JWT Error:", err.Error())
	}

	// When you use jwt.New(), the function is already automatically called for checking,
	// which means you don't need to call it again.
	errInit := authMiddleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:", errInit.Error())
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
			userAuth.GET("user/info", api.UserInfo)
			userAuth.DELETE("user/logout", authMiddleware.LogoutHandler)
		}
	}
	return r
}
