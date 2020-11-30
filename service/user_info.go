package service

import (
	"call-up/model"
	"call-up/serializer"
	"github.com/gin-gonic/gin"
)

// UserInfo 用户信息服务
type UserInfo struct {
	UserID string `form:"user" json:"user" binding:"required,min=4,max=16"`
}

// Info 用户信息函数
func (service *UserInfo) Info(c *gin.Context) serializer.Response {
	curUser, _ := c.Get("user")

	if service.UserID != curUser.(*model.User).UserID {
		return serializer.Err(serializer.CodeParamErr, "无权限", nil)
	}

	var user model.User
	if err := model.DB.Where("user_id = ?", service.UserID).First(&user).Error; err != nil {
		return serializer.Err(serializer.CodeParamErr, "用户不存在", nil)
	}

	resp := serializer.BuildUserInfoResponse(user)
	resp.Msg = "查询成功"
	return resp
}
