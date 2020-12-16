package service

import (
	"call-up/model"
	"call-up/serializer"

	"github.com/gin-gonic/gin"
)

// CallupList 召集令列表服务
type CallupList struct {
	Offset int `form:"offset" json:"offset"`
	Limit  int `form:"limit" json:"limit"`
}

// List 召集令列表
func (service *CallupList) List(c *gin.Context) serializer.Response {
	if service.Limit == 0 {
		service.Limit = 10
	}

	var total int64 = 0
	if err := model.DB.Model(&model.Callup{}).Count(&total).Error; err != nil {
		return serializer.Err(serializer.CodeDBError, "接令请求列表查询失败", err)
	}

	var callups []model.Callup
	if err := model.DB.Model(&model.Callup{}).Limit(service.Limit).Offset(service.Offset).Find(&callups).Error; err != nil {
		return serializer.Err(serializer.CodeDBError, "接令请求列表查询失败", err)
	}

	resp := serializer.BuildCallupListResponse(callups, total)
	resp.Msg = "查询成功"
	return resp
}
