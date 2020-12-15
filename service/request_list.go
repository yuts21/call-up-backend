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
	user, _ := c.Get("user")
	requester := user.(*model.User)

	if service.Limit == 0 {
		service.Limit = 10
	}

	total := model.DB.Model(&requester).Association("Request").Count()

	var results []serializer.RequestListItem
	if err := model.DB.Model(&model.Request{}).Select("requests.id as request_id, callups.name as callup_name, requests.status as status").
		Joins("join callups on requests.callup_id = callups.id").
		Where("requester_id = ?", requester.ID).
		Limit(service.Limit).
		Offset(service.Offset).
		Scan(&results).Error; err != nil {
		return serializer.Err(serializer.CodeDBError, "接令请求列表查询失败", err)
	}

	resp := serializer.BuildListResponse(results, total)
	resp.Msg = "查询成功"
	return resp
}
