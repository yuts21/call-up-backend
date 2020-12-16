package service

import (
	"call-up/model"
	"call-up/serializer"

	"github.com/gin-gonic/gin"
)

// CallupRequestList 召集令请求列表服务
type CallupRequestList struct {
	ID     uint `form:"id" json:"id" binding:"required"`
	Offset int  `form:"offset" json:"offset"`
	Limit  int  `form:"limit" json:"limit"`
}

// List 召集令请求列表
func (service *CallupRequestList) List(c *gin.Context) serializer.Response {
	curUser, _ := c.Get("user")
	user := curUser.(*model.User)

	var callup model.Callup
	if err := model.DB.Where("id = ? and sponsor_id = ?", service.ID, user.ID).First(&callup).Error; err != nil {
		return serializer.Err(serializer.CodeDBError, "召集令查询失败", err)
	}

	if service.Limit == 0 {
		service.Limit = 10
	}

	total := model.DB.Model(&callup).Association("Request").Count()

	var results []serializer.CallupRequestListItem
	if err := model.DB.Model(&model.Request{}).Select("requests.id as id, requests.requester_id as requester_id, users.name as requester_name, requests.status as status").
		Joins("join users on requests.requester_id = users.id").
		Where("callup_id = ?", callup.ID).
		Limit(service.Limit).
		Offset(service.Offset).
		Scan(&results).Error; err != nil {
		return serializer.Err(serializer.CodeDBError, "接令请求列表查询失败", err)
	}

	resp := serializer.BuildListResponse(results, total)
	resp.Msg = "查询成功"
	return resp
}
