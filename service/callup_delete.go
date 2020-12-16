package service

import (
	"call-up/model"
	"call-up/serializer"

	"github.com/gin-gonic/gin"
)

// CallupDelete 删除召集令服务
type CallupDelete struct {
	ID uint `form:"id" json:"id" binding:"required"`
}

// Delete 删除召集令函数
func (service *CallupDelete) Delete(c *gin.Context) serializer.Response {
	curUser, _ := c.Get("user")
	user := curUser.(*model.User)

	var callup model.Callup
	if err := model.DB.Where("id = ? and sponsor_id = ?", service.ID, user.ID).First(&callup).Error; err != nil {
		return serializer.Err(serializer.CodeDBError, "召集令查询失败", err)
	}

	if !callup.Canceled {
		return serializer.Err(serializer.CodeParamErr, "该召集令未被取消", nil)
	}

	if err := model.DB.Delete(&callup).Error; err != nil {
		return serializer.Err(serializer.CodeDBError, "删除召集令失败", err)
	}

	return serializer.Success("删除召集令成功")
}
