package service

import (
	"call-up/model"
	"call-up/serializer"

	"github.com/gin-gonic/gin"
)

// CallupMine 自己发布的召集令列表服务
type CallupMine struct {
	Offset int `form:"offset" json:"offset"`
	Limit  int `form:"limit" json:"limit"`
}

// List 召集令列表
func (service *CallupMine) List(c *gin.Context) serializer.Response {
	curUser, _ := c.Get("user")
	user := curUser.(*model.User)

	if service.Limit == 0 {
		service.Limit = 10
	}

	total := model.DB.Model(&user).Association("Callup").Count()

	var callups []model.Callup
	if err := model.DB.Where("sponsor_id = ?", user.ID).Limit(service.Limit).Offset(service.Offset).Find(&callups).Error; err != nil {
		return serializer.Err(serializer.CodeDBError, "召集令列表查询失败", err)
	}

	resp := serializer.BuildCallupListResponse(callups, total)
	resp.Msg = "查询成功"
	return resp
}
