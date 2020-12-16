package service

import (
	"call-up/model"
	"call-up/serializer"

	"github.com/gin-gonic/gin"
)

// CallupMine 自己发布的召集令列表服务
type CallupMine struct {
	Type   *uint8  `form:"type" json:"type" binding:"omitempty,gt=0"`
	Name   *string `form:"name" json:"name"`
	Offset int     `form:"offset" json:"offset"`
	Limit  int     `form:"limit" json:"limit"`
}

// List 召集令列表
func (service *CallupMine) List(c *gin.Context) serializer.Response {
	curUser, _ := c.Get("user")
	user := curUser.(*model.User)

	if service.Limit == 0 {
		service.Limit = 10
	}

	db := model.DB.Model(&model.Callup{}).Where("sponsor_id = ?", user.ID)
	if service.Type != nil {
		db = db.Where("type = ?", *service.Type)
	}
	if service.Name != nil {
		db = db.Where("name like ?", "%"+*service.Name+"%")
	}

	var total int64 = 0
	if err := db.Count(&total).Error; err != nil {
		return serializer.Err(serializer.CodeDBError, "召集令列表查询失败", err)
	}

	var callups []model.Callup
	if err := db.Limit(service.Limit).Offset(service.Offset).Find(&callups).Error; err != nil {
		return serializer.Err(serializer.CodeDBError, "召集令列表查询失败", err)
	}

	resp := serializer.BuildCallupListResponse(callups, total)
	resp.Msg = "查询成功"
	return resp
}
