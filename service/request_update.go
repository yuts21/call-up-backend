package service

import (
	"call-up/model"
	"call-up/serializer"

	"github.com/gin-gonic/gin"
)

// RequestUpdate 接令请求修改服务
type RequestUpdate struct {
	RequestID   uint   `form:"request_id" json:"request_id" binding:"required"`
	Description string `form:"description" json:"description" binding:"required"`
}

// Update 删除接令请求
func (service *RequestUpdate) Update(c *gin.Context) serializer.Response {
	curUser, _ := c.Get("user")
	user := curUser.(*model.User)

	var request model.Request

	if err := model.DB.Where("id = ?", service.RequestID).First(&request).Error; err != nil {
		return serializer.Err(serializer.CodeDBError, "接令请求不存在", err)
	}
	if !user.Type && user.ID != request.RequesterID {
		return serializer.Err(serializer.CodeNoRightErr, "无权限", nil)
	}
	if err := model.DB.
		Model(&model.Request{}).
		Where("id = ? AND status != ?", request.ID, model.Agreed).
		Update("description", service.Description).Error; err != nil {
		return serializer.Err(serializer.CodeDBError, "接令请求修改失败", err)
	}

	resp := serializer.Success("修改成功")
	return resp
}
