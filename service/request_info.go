package service

import (
	"call-up/model"
	"call-up/serializer"

	"github.com/gin-gonic/gin"
)

// RequestInfo 接令请求查询服务
type RequestInfo struct {
	RequestID uint `form:"request_id" json:"request_id" binding:"required"`
}

// Info 查询接令请求
func (service *RequestInfo) Info(c *gin.Context) serializer.Response {
	var request model.Request
	user, _ := c.Get("user")
	requester := user.(*model.User)

	if err := model.DB.
		Where("id = ? and requester_id = ?", service.RequestID, requester.ID).
		First(&request).Error; err != nil {
		return serializer.Err(serializer.CodeDBError, "接令请求查询失败", err)
	}

	resp := serializer.BuildRequestInfoResponse(request)
	resp.Msg = "查询成功"
	return resp
}
