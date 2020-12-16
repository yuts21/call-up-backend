package service

import (
	"call-up/model"
	"call-up/serializer"

	"github.com/gin-gonic/gin"
)

// CallupInfo 召集令信息服务
type CallupInfo struct {
	ID uint `form:"id" json:"id" binding:"required"`
}

// Info 召集令信息函数
func (service *CallupInfo) Info(c *gin.Context) serializer.Response {
	var callup model.Callup
	if err := model.DB.Where("id = ?", service.ID).First(&callup).Error; err != nil {
		return serializer.Err(serializer.CodeDBError, "召集令查询失败", err)
	}

	var sponsor model.User
	if err := model.DB.Model(&callup).Association("Sponsor").Find(&sponsor); err != nil {
		return serializer.Err(serializer.CodeDBError, "令主查询失败", err)
	}

	resp := serializer.BuildCallupInfoResponse(callup, sponsor)
	resp.Msg = "查询成功"
	return resp
}
