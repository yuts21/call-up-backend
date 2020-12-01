package service

import (
	"call-up/model"
	"call-up/serializer"

	"github.com/gin-gonic/gin"
)

// RequestList 接令请求列表服务
type RequestList struct {
}

// List 接令请求列表
func (service *RequestList) List(c *gin.Context) serializer.Response {
	requests := []model.Request{}
	user, _ := c.Get("user")
	requester := user.(*model.User)

	if err := model.DB.
		Where("requester_id = ?", requester.ID).
		Find(&requests).Error; err != nil {
		return serializer.Err(serializer.CodeDBError, "接令请求列表查询失败", err)
	}

	resp := serializer.BuildRequestListResponse(requests)
	resp.Msg = "查询成功"
	return resp
}
