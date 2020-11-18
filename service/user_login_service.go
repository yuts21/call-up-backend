package service

import (
	"call-up/model"
	"call-up/serializer"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserLoginService 用户登录服务
type UserLoginService struct {
	UserID string `form:"user_id" json:"user_id" binding:"required,min=4,max=16"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=16"`
}

// setSession 设置session
func (service *UserLoginService) setSession(c *gin.Context, user model.User) {
	s := sessions.Default(c)
	s.Clear()
	s.Set("user_id", user.UserID)
	s.Save()
}

// Login 用户登录函数
func (service *UserLoginService) Login(c *gin.Context) serializer.Response {
	var user model.User
	if err := model.DB.Where("user_id = ?", service.UserID).First(&user).Error; err != nil {
		return serializer.Err(serializer.CodeParamErr, "用户名或密码错误", nil)
	}

	if ! user.CheckPassword(service.Password) {
		return serializer.Err(serializer.CodeParamErr, "用户名或密码错误", nil)
	}

	// 设置session
	service.setSession(c, user)

	resp := serializer.BuildUserResponse(user)
	resp.Msg = "登录成功"
	return resp
}
