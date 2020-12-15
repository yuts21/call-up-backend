package service

import (
	"call-up/model"
	"call-up/serializer"

	"github.com/gin-gonic/gin"
)

// CallupRequestHandle 召集令请求处理服务
type CallupRequestHandle struct {
	RequestID   uint  `form:"request_id" json:"request_id" binding:"required"`
	Instruction uint8 `form:"inst" json:"inst" binding:"required,gte=1,lte=2"`
}

// Handle 处理接令请求
func (service *CallupRequestHandle) Handle(c *gin.Context) serializer.Response {
	curUser, _ := c.Get("user")
	user := curUser.(*model.User)

	var request model.Request
	if err := model.DB.Where("id = ?", service.RequestID).First(&request).Error; err != nil {
		return serializer.Err(serializer.CodeDBError, "接令请求查询失败", err)
	}

	var callup model.Callup
	if err := model.DB.Model(&request).Association("Callup").Find(&callup); err != nil {
		return serializer.Err(serializer.CodeDBError, "召集令查询失败", err)
	}

	if callup.SponsorID != user.ID {
		return serializer.Err(serializer.CodeNoRightErr, "无权限", nil)
	}

	status := callup.Status()
	if status == model.Expired || status == model.Canceled {
		return serializer.Err(serializer.CodeParamErr, "召集令已逾期", nil)
	}

	if service.Instruction == 2 {
		if err := model.DB.Model(&request).Update("status", model.Denied).Error; err != nil {
			return serializer.Err(serializer.CodeDBError, "操作失败", err)
		}
		return serializer.Success("操作成功")
	}

	if status == model.Completed {
		return serializer.Err(serializer.CodeParamErr, "召集令人数已满", nil)
	}

	if err := model.DB.Model(&request).Update("status", model.Agreed).Error; err != nil {
		return serializer.Err(serializer.CodeDBError, "操作失败", err)
	}

	return serializer.Success("操作成功")
}
