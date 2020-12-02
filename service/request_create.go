package service

import (
	"call-up/model"
	"call-up/serializer"

	"github.com/gin-gonic/gin"
)

// RequestCreate 接令请求创建服务
type RequestCreate struct {
	CallupID    uint   `form:"callup_id" json:"callup_id" binding:"required"`
	Description string `form:"description" json:"description"`
}

// Create 创建接令请求
func (service *RequestCreate) Create(c *gin.Context) serializer.Response {
	user, _ := c.Get("user")
	Requester := user.(*model.User)

	request := model.Request{
		CallupID:    service.CallupID,
		RequesterID: Requester.ID,
		Description: service.Description,
		Status:      model.Unprocessed,
	}

	if err := model.DB.Create(&request).Error; err != nil {
		return serializer.Err(serializer.CodeDBError, "接令请求创建失败", err)
	}

	return serializer.Success("接令请求创建成功")
}
