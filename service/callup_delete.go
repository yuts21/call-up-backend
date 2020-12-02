package service

import (
	"call-up/model"
	"call-up/serializer"
	"github.com/gin-gonic/gin"
)

// CallupDelete 删除召集令服务
type CallupDelete struct {
	ID uint `form:"ID" json:"ID" binding:"required"`
}

// Delete 删除召集令函数
func (service *CallupDelete) Delete(c *gin.Context) serializer.Response {
	curUser, _ := c.Get("user")
	user := curUser.(*model.User)

	var callup model.Callup
	if err := model.DB.Where("id = ? and lord_id = ?", service.ID, user.ID).First(&callup).Error; err != nil {
		return serializer.Err(serializer.CodeDBError, "召集令查询失败", err)
	}

	count := 0
	var requests []model.Request
	model.DB.Model(&callup).Related(&requests).Count(&count)
	if count > 0 {
		return serializer.Err(serializer.CodeParamErr, "该召集令已有响应者", nil)
	}

	if err := model.DB.Delete(&callup).Error; err != nil {
		return serializer.Err(serializer.CodeDBError, "删除召集令失败", err)
	}

	return serializer.Success("删除召集令成功")
}
