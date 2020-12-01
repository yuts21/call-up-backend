package service

import (
	"call-up/model"
	"call-up/serializer"

	"github.com/gin-gonic/gin"
)

// RequestShow 接令请求查询服务
type RequestShow struct {
	RequestID uint `form:"request_id" json:"request_id" binding:"required"`
}

// Show 查询接令请求
func (service *RequestShow) Show(c *gin.Context) serializer.Response {
	var request model.Request
	user, _ := c.Get("user")
	requester := user.(*model.User)

	if err := model.DB.
		Where("id = ? AND requester_id = ?", service.RequestID, requester.ID).
		First(&request).Error; err != nil {
		return serializer.Err(serializer.CodeDBError, "接令请求查询失败", err)
	}

	resp := serializer.BuildRequestShowResponse(request)
	resp.Msg = "接令请求查询成功"
	return resp
}
