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

// UserPasswordUpdate 修改密码
func UserPasswordUpdate(c *gin.Context) {
	var serv service.UserPasswordUpdate
	if err := c.ShouldBind(&serv); err == nil {
		res := serv.Update(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UserInfoUpdate 修改用户信息
func UserInfoUpdate(c *gin.Context) {
	var serv service.UserInfoUpdate
	if err := c.ShouldBind(&serv); err == nil {
		res := serv.Update(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
