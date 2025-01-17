package service

import (
	"call-up/model"
	"call-up/serializer"

	"github.com/gin-gonic/gin"
)

// RequestInfo 接令请求查询服务
type RequestInfo struct {
	ID uint `form:"id" json:"id" binding:"required"`
}

// Info 查询接令请求
func (service *RequestInfo) Info(c *gin.Context) serializer.Response {
	curUser, _ := c.Get("user")
	user := curUser.(*model.User)

	var request model.Request
	if err := model.DB.Where("id = ? and requester_id = ?", service.ID, user.ID).First(&request).Error; err != nil {
		return serializer.Err(serializer.CodeDBError, "接令请求查询失败", err)
	}

	var callup model.Callup
	if err := model.DB.Model(&request).Association("Callup").Find(&callup); err != nil {
		return serializer.Err(serializer.CodeDBError, "召集令查询失败", err)
	}

	resp := serializer.BuildRequestInfoResponse(request, callup)
	resp.Msg = "查询成功"
	return resp
}
