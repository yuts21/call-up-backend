package service

import (
	"call-up/model"
	"call-up/serializer"

	"github.com/gin-gonic/gin"
)

// RequestCreate 接令请求创建服务
type RequestCreate struct {
	CallupID    uint   `form:"callup_id" json:"callup_id" binding:"required"`
	Description string `form:"descrpt" json:"descrpt"`
}

// Create 创建接令请求
func (service *RequestCreate) Create(c *gin.Context) serializer.Response {
	user, _ := c.Get("user")
	Requester := user.(*model.User)

	var callup model.Callup
	if err := model.DB.Where("id = ?", service.CallupID).First(&callup).Error; err != nil {
		return serializer.Err(serializer.CodeDBError, "召集令查询失败", err)
	}

	if callup.Status() != model.Waiting {
		return serializer.Err(serializer.CodeParamErr, "召集令不处于待响应状态", nil)
	}

	request := model.Request{
		CallupID:    service.CallupID,
		RequesterID: Requester.ID,
		Description: service.Description,
		Status:      model.Unprocessed,
	}

	if err := model.DB.Create(&request).Error; err != nil {
		return serializer.Err(serializer.CodeDBError, "接令请求创建失败", err)
	}

	return serializer.Success("创建接令请求成功")
}
