package service

import (
	"call-up/model"
	"call-up/serializer"

	"github.com/gin-gonic/gin"
)

// RequestDelete 接令请求删除服务
type RequestDelete struct {
	RequestID uint `form:"request_id" json:"request_id" binding:"required"`
}

// Delete 删除接令请求
func (service *RequestDelete) Delete(c *gin.Context) serializer.Response {
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
		Where("id = ? AND status != ?", request.ID, model.Agreed).
		Delete(&model.Request{}).Error; err != nil {
		return serializer.Err(serializer.CodeDBError, "接令请求删除失败", err)
	}

	resp := serializer.Success("删除成功")
	return resp
}
