package service

import (
	"call-up/model"
	"call-up/serializer"

	"github.com/gin-gonic/gin"
)

// CallupCancel 取消召集令服务
type CallupCancel struct {
	ID uint `form:"ID" json:"ID" binding:"required"`
}

// Cancel 取消召集令函数
func (service *CallupCancel) Cancel(c *gin.Context) serializer.Response {
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

	if err := model.DB.Model(&callup).Update("canceled", true).Error; err != nil {
		return serializer.Err(serializer.CodeDBError, "取消召集令失败", err)
	}

	return serializer.Success("取消召集令成功")
}
