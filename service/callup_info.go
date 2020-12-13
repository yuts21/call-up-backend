package service

import (
	"call-up/model"
	"call-up/serializer"

	"github.com/gin-gonic/gin"
)

// CallupInfo 召集令信息服务
type CallupInfo struct {
	ID uint `form:"ID" json:"ID" binding:"required"`
}

// Info 召集令信息函数
func (service *CallupInfo) Info(c *gin.Context) serializer.Response {
	curUser, _ := c.Get("user")
	user := curUser.(*model.User)

	var callup model.Callup
	if err := model.DB.Where("id = ? and lord_id = ?", service.ID, user.ID).First(&callup).Error; err != nil {
		return serializer.Err(serializer.CodeDBError, "召集令查询失败", err)
	}

	resp := serializer.BuildCallupInfoResponse(callup)
	resp.Msg = "查询成功"
	return resp
}
