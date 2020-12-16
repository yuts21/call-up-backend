package service

import (
	"call-up/model"
	"call-up/serializer"
	"time"

	"github.com/gin-gonic/gin"
)

// CallupUpdate 修改召集令服务
type CallupUpdate struct {
	ID          uint    `form:"id" json:"id" binding:"required"`
	Type        *uint8  `form:"type" json:"type"`
	Name        *string `form:"name" json:"name"`
	Description *string `form:"descrpt" json:"descrpt"`
	Capacity    *uint   `form:"cap" json:"cap" binding:"omitempty,gt=0"`
	EndDate     *int64  `form:"end_date" json:"end_date"`
}

// Update 修改召集令函数
func (service *CallupUpdate) Update(c *gin.Context) serializer.Response {
	curUser, _ := c.Get("user")
	user := curUser.(*model.User)

	var callup model.Callup
	if err := model.DB.Where("id = ? and sponsor_id = ?", service.ID, user.ID).First(&callup).Error; err != nil {
		return serializer.Err(serializer.CodeDBError, "召集令查询失败", err)
	}

	count := model.DB.Model(&callup).Association("Request").Count()
	if count > 0 {
		return serializer.Err(serializer.CodeParamErr, "该召集令已有响应者", nil)
	}

	// 修改数据库
	callupNew := make(map[string]interface{})
	if service.Type != nil {
		callupNew["type"] = *service.Type
	}
	if service.Name != nil {
		callupNew["name"] = *service.Name
	}
	if service.Description != nil {
		callupNew["description"] = *service.Description
	}
	if service.Capacity != nil {
		callupNew["capacity"] = *service.Capacity
	}
	if service.EndDate != nil {
		callupNew["end_date"] = time.Unix(*service.EndDate, 0)
	}
	if err := model.DB.Model(&callup).Updates(callupNew).Error; err != nil {
		return serializer.Err(serializer.CodeDBError, "修改召集令失败", err)
	}

	return serializer.Success("修改召集令成功")
}
