package service

import (
	"call-up/model"
	"call-up/serializer"

	"github.com/gin-gonic/gin"
)

// UserInfo 用户信息服务
type UserInfo struct {
	ID uint `form:"id" json:"id" binding:"required"`
}

// Info 用户信息函数
func (service *UserInfo) Info(c *gin.Context) serializer.Response {
	curUser, _ := c.Get("user")
	user := *curUser.(*model.User)
	if user.ID == service.ID {
		resp := serializer.BuildUserInfoResponse(user)
		resp.Msg = "查询成功"
		return resp
	}

	if user.Type {
		user, err := model.GetUser(service.ID)
		if err != nil {
			return serializer.Err(serializer.CodeDBError, "用户查询失败", err)
		}
		resp := serializer.BuildUserInfoResponse(user)
		resp.Msg = "查询成功"
		return resp
	}

	var count int64 = 0
	if err := model.DB.Model(&model.Request{}).
		Joins("join callups on requests.callup_id = callups.id").
		Where("callups.sponsor_id = ? and requests.requester_id = ?", user.ID, service.ID).
		Count(&count).Error; err != nil {
		return serializer.Err(serializer.CodeDBError, "请求信息查询失败", err)
	}
	if count == 0 {
		return serializer.Err(serializer.CodeNoRightErr, "无权限", nil)
	}

	user, err := model.GetUser(service.ID)
	if err != nil {
		return serializer.Err(serializer.CodeDBError, "用户查询失败", err)
	}
	resp := serializer.BuildUserInfoResponse(user)
	resp.Msg = "查询成功"
	return resp
}
