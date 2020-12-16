package service

import (
	"call-up/model"
	"call-up/serializer"

	"github.com/gin-gonic/gin"
)

// CallupRequestInfo 召集令请求查询服务
type CallupRequestInfo struct {
	RequestID uint `form:"request_id" json:"request_id" binding:"required"`
}

// Info 查询接令请求
func (service *CallupRequestInfo) Info(c *gin.Context) serializer.Response {
	curUser, _ := c.Get("user")
	user := curUser.(*model.User)

	var request model.Request
	if err := model.DB.Where("id = ?", service.RequestID).First(&request).Error; err != nil {
		return serializer.Err(serializer.CodeDBError, "接令请求查询失败", err)
	}

	var callup model.Callup
	if err := model.DB.Model(&request).Association("Callup").Find(&callup); err != nil {
		return serializer.Err(serializer.CodeDBError, "召集令查询失败", err)
	}

	if callup.SponsorID != user.ID {
		return serializer.Err(serializer.CodeNoRightErr, "无权限", nil)
	}

	var requester model.User
	if err := model.DB.Model(&request).Association("Requester").Find(&requester); err != nil {
		return serializer.Err(serializer.CodeDBError, "请求者查询失败", nil)
	}

	resp := serializer.BuildCallupRequestInfoResponse(request, requester)
	resp.Msg = "查询成功"
	return resp
}
