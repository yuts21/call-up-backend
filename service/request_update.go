package service

import (
	"call-up/model"
	"call-up/serializer"

	"github.com/gin-gonic/gin"
)

// RequestUpdate 接令请求修改服务
type RequestUpdate struct {
	ID   uint   `form:"ID" json:"ID" binding:"required"`
	Description string `form:"descrpt" json:"descrpt"`
}

// Update 修改接令请求
func (service *RequestUpdate) Update(c *gin.Context) serializer.Response {
	curUser, _ := c.Get("user")
	user := curUser.(*model.User)

	var request model.Request
	if err := model.DB.Where("id = ? and requester_id = ?", service.ID, user.ID).First(&request).Error; err != nil {
		return serializer.Err(serializer.CodeDBError, "接令请求查询失败", err)
	}

	if err := model.DB.
		Model(&model.Request{}).
		Where("id = ? and status != ?", request.ID, model.Agreed).
		Update("description", service.Description).Error; err != nil {
		return serializer.Err(serializer.CodeDBError, "接令请求修改失败", err)
	}

	resp := serializer.Success("修改成功")
	return resp
}
