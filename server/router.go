package server

import (
	"call-up/api"
	"call-up/middleware"
	"github.com/gin-gonic/gin"
	"os"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())

	// 路由
	v1 := r.Group("/api")
	{
		v1.POST("ping", api.Ping)

		// 用户注册、登录
		v1.POST("user/login", api.UserLogin)
		v1.POST("user/reg", api.UserRegister)

		// 需要用户登录保护的
		userAuth := v1.Group("")
		{
			userAuth.Use(middleware.CurrentUser())
			userAuth.Use(middleware.UserAuthRequired())

			userAuth.GET("user/info", api.UserInfo)
			userAuth.DELETE("user/logout", api.UserLogout)
		}
	}
	return r
}
