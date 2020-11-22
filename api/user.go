package api

import (
	"call-up/service"
	"github.com/gin-gonic/gin"
)

// UserRegister 用户注册接口
func UserRegister(c *gin.Context) {
	var serv service.UserRegister
	if err := c.ShouldBind(&serv); err == nil {
		res := serv.Register()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}


// UserInfo 用户信息
func UserInfo(c *gin.Context) {
	var serv service.UserInfo
	if err := c.ShouldBind(&serv); err == nil {
		res := serv.Info(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
