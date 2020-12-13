package service

import (
	"call-up/model"
	"call-up/serializer"

	"github.com/gin-gonic/gin"
)

// RequestList 接令请求列表服务
type RequestList struct {
	Offset int `form:"offset" json:"offset"`
	Limit  int `form:"limit" json:"limit"`
}

// List 接令请求列表
func (service *RequestList) List(c *gin.Context) serializer.Response {
	requests := []model.Request{}
	user, _ := c.Get("user")
	requester := user.(*model.User)

	if service.Limit == 0 {
		service.Limit = 10
	}

	total := 0
	if err := model.DB.Model(&model.Request{}).Where("requester_id = ?", requester.ID).Count(&total).Error; err != nil {
		return serializer.Err(serializer.CodeDBError, "接令请求列表查询失败", err)
	}

	if err := model.DB.
		Where("requester_id = ?", requester.ID).
		Limit(service.Limit).
		Offset(service.Offset).
		Find(&requests).Error; err != nil {
		return serializer.Err(serializer.CodeDBError, "接令请求列表查询失败", err)
	}

	resp := serializer.BuildListResponse(serializer.BuildRequestList(requests), uint(total))
	resp.Msg = "查询成功"
	return resp
}
