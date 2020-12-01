package service

import (
	"call-up/model"
	"call-up/serializer"
	"github.com/gin-gonic/gin"
	"time"
)

// CallupCreate 创建召集令服务
type CallupCreate struct {
	Type        uint8  `form:"type" json:"type" binding:"required"`
	Name        string `form:"name" json:"name" binding:"required"`
	Description string `form:"descrpt" json:"descrpt"`
	Capacity    uint   `form:"cap" json:"cap" binding:"required,gt=0"`
	EndDate     int64  `form:"end_date" json:"end_date" binding:"required"`
	Picture     []byte `form:"picture" json:"picture"`
}

// Create 创建召集令
func (service *CallupCreate) Create(c *gin.Context) serializer.Response {
	curUser, _ := c.Get("user")
	Lord := curUser.(*model.User)

	callup := model.Callup{
		LordID: Lord.ID,
		Type: service.Type,
		Name: service.Name,
		Description: service.Description,
		Capacity: service.Capacity,
		EndDate: time.Unix(service.EndDate, 0),
		Picture: service.Picture,
		Status: model.Waiting,
	}

	if err := model.DB.Create(&callup).Error; err != nil {
		return serializer.Err(serializer.CodeDBError, "召集令创建失败", err)
	}

	resp := serializer.Success("召集令创建成功")
	return resp
}
