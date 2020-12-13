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
	if err := model.DB.Where("id = ? and lord_id = ?", service.ID, user.ID).First(&callup).Error; err != nil {
		return serializer.Err(serializer.CodeDBError, "召集令查询失败", err)
	}

	count := 0
	if err := model.DB.Model(&model.Request{}).Where("callup_id = ?", callup.ID).Count(&count).Error; err != nil {
		return serializer.Err(serializer.CodeDBError, "召集令接令请求查询失败", err)
	}

	if count > 0 {
		return serializer.Err(serializer.CodeParamErr, "该召集令已有响应者", nil)
	}

	if err := model.DB.Model(&callup).Update("canceled", true).Error; err != nil {
		return serializer.Err(serializer.CodeDBError, "取消召集令失败", err)
	}

	return serializer.Success("取消召集令成功")
}
