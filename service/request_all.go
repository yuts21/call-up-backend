package service

import (
	"call-up/model"
	"call-up/serializer"

	"github.com/gin-gonic/gin"
)

// RequestAll 全部接令请求列表服务
type RequestAll struct {
	Offset int `form:"offset" json:"offset"`
	Limit  int `form:"limit" json:"limit"`
}

// List 接令请求列表
func (service *RequestAll) List(c *gin.Context) serializer.Response {
	if service.Limit == 0 {
		service.Limit = 10
	}

	var total int64 = 0
	if err := model.DB.Model(&model.Request{}).Count(&total).Error; err != nil {
		return serializer.Err(serializer.CodeDBError, "接令请求列表查询失败", err)
	}

	var results []serializer.RequestAllItem
	if err := model.DB.Model(&model.Request{}).Select(`requests.id as id, requests.callup_id as callup_id, callups.name as callup_name, 
		requests.requester_id as requester_id, users.name as requester_name, requests.status as status`).
		Joins("join callups on requests.callup_id = callups.id").
		Joins("join users on requests.requester_id = users.id").
		Limit(service.Limit).
		Offset(service.Offset).
		Scan(&results).Error; err != nil {
		return serializer.Err(serializer.CodeDBError, "接令请求列表查询失败", err)
	}

	resp := serializer.BuildListResponse(results, total)
	resp.Msg = "查询成功"
	return resp
}
