package service

import (
	"call-up/model"
	"call-up/serializer"

	"github.com/gin-gonic/gin"
)

// RequestCancel 接令请求取消服务
type RequestCancel struct {
	ID uint `form:"id" json:"id" binding:"required"`
}

// Cancel 取消接令请求
func (service *RequestCancel) Cancel(c *gin.Context) serializer.Response {
	curUser, _ := c.Get("user")
	user := curUser.(*model.User)

	var request model.Request
	if err := model.DB.Where("id = ? and requester_id = ?", service.ID, user.ID).First(&request).Error; err != nil {
		return serializer.Err(serializer.CodeDBError, "接令请求查询失败", err)
	}

	if request.Status != model.Unprocessed {
		return serializer.Err(serializer.CodeParamErr, "接令请求不处于未处理状态", nil)
	}

	if err := model.DB.Model(&request).Update("status", model.Abolished).Error; err != nil {
		return serializer.Err(serializer.CodeDBError, "接令请求取消失败", err)
	}

	resp := serializer.Success("取消成功")
	return resp
}
