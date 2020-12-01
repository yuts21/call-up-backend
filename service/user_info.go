package service

import (
	"call-up/model"
	"call-up/serializer"
	"github.com/gin-gonic/gin"
)

// UserInfo 用户信息服务
type UserInfo struct {
	ID uint `form:"UID" json:"UID" binding:"required"`
}

// Info 用户信息函数
func (service *UserInfo) Info(c *gin.Context) serializer.Response {
	curUser, _ := c.Get("user")

	if service.ID != curUser.(*model.User).ID {
		return serializer.Err(serializer.CodeNoRightErr, "无权限", nil)
	}

	var user model.User
	if err := model.DB.Where("id = ?", service.ID).First(&user).Error; err != nil {
		return serializer.Err(serializer.CodeParamErr, "用户不存在", nil)
	}

	resp := serializer.BuildUserInfoResponse(user)
	resp.Msg = "查询成功"
	return resp
}
