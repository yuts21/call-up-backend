package service

import (
	"call-up/model"
	"call-up/serializer"

	"github.com/gin-gonic/gin"
)

// RequestDelete 接令请求删除服务
type RequestDelete struct {
	ID uint `form:"ID" json:"ID binding:"required"`
}

// Delete 删除接令请求
func (service *RequestDelete) Delete(c *gin.Context) serializer.Response {
	curUser, _ := c.Get("user")
	user := curUser.(*model.User)

	var request model.Request
	if err := model.DB.Where("id = ? and requester_id = ?", service.ID, user.ID).First(&request).Error; err != nil {
		return serializer.Err(serializer.CodeDBError, "接令请求查询失败", err)
	}

	if err := model.DB.
		Where("id = ? and status != ?", request.ID, model.Agreed).
		Delete(&model.Request{}).Error; err != nil {
		return serializer.Err(serializer.CodeDBError, "接令请求删除失败", err)
	}

	resp := serializer.Success("删除成功")
	return resp
}
