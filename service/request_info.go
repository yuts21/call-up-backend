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
	curUser, _ := c.Get("user")
	user := curUser.(*model.User)

	var request model.Request
	if err := model.DB.Where("id = ?", service.RequestID).First(&request).Error; err != nil {
		return serializer.Err(serializer.CodeDBError, "接令请求查询失败", err)
	}
	if !user.Type && user.ID != request.RequesterID {
		return serializer.Err(serializer.CodeNoRightErr, "无权限", nil)
	}

	resp := serializer.BuildRequestInfoResponse(request)
	resp.Msg = "查询成功"
	return resp
}
