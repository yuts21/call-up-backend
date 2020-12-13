package service

import (
	"call-up/model"
	"call-up/serializer"

	"github.com/gin-gonic/gin"
)

// RequestInfo 接令请求查询服务
type RequestInfo struct {
	ID uint `form:"ID" json:"ID" binding:"required"`
}

// Info 查询接令请求
func (service *RequestInfo) Info(c *gin.Context) serializer.Response {
	curUser, _ := c.Get("user")
	user := curUser.(*model.User)

	var request model.Request
	if err := model.DB.Where("id = ? and requester_id = ?", service.ID, user.ID).First(&request).Error; err != nil {
		return serializer.Err(serializer.CodeDBError, "接令请求查询失败", err)
	}

	resp := serializer.BuildRequestInfoResponse(request)
	resp.Msg = "查询成功"
	return resp
}
