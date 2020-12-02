package service

import (
	"call-up/model"
	"call-up/serializer"

	"github.com/gin-gonic/gin"
)

// UserInfo 用户信息服务
type UserInfo struct {
}

// Info 用户信息函数
func (service *UserInfo) Info(c *gin.Context) serializer.Response {
	curUser, _ := c.Get("user")
	user := *curUser.(*model.User)

	resp := serializer.BuildUserInfoResponse(user)
	resp.Msg = "查询成功"
	return resp
}
