package service

import (
	"call-up/model"
	"call-up/serializer"

	"github.com/gin-gonic/gin"
)

// RequestCallup 已接令的召集令列表服务
type RequestCallup struct {
	Offset int `form:"offset" json:"offset"`
	Limit  int `form:"limit" json:"limit"`
}

// List 召集令列表
func (service *RequestCallup) List(c *gin.Context) serializer.Response {
	curUser, _ := c.Get("user")
	user := curUser.(*model.User)

	if service.Limit == 0 {
		service.Limit = 10
	}

	var total int64 = 0
	if err := model.DB.Model(&model.Callup{}).
		Joins("join requests on callups.id = requests.callup_id").
		Where("requests.requester_id = ? and requests.status = ?", user.ID, model.Agreed).
		Count(&total).Error; err != nil {
			return serializer.Err(serializer.CodeDBError, "召集令列表查询失败", err)
	}

	var callups []model.Callup
	if err := model.DB.Model(&model.Callup{}).Select("callups.*").
		Joins("join requests on callups.id = requests.callup_id").
		Where("requests.requester_id = ? and requests.status = ?", user.ID, model.Agreed).
		Limit(service.Limit).Offset(service.Offset).Scan(&callups).Error; err != nil {
			return serializer.Err(serializer.CodeDBError, "召集令列表查询失败", err)
	}

	resp := serializer.BuildCallupListResponse(callups, total)
	resp.Msg = "查询成功"
	return resp
}
