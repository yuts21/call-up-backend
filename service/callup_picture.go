package service

import (
	"call-up/model"
	"call-up/serializer"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CallupPicture 召集令图片服务
type CallupPicture struct {
	ID uint `form:"ID" json:"ID" binding:"required"`
}

// GetPicture 召集令图片函数
func (service *CallupPicture) GetPicture(c *gin.Context) {
	curUser, _ := c.Get("user")
	user := curUser.(*model.User)

	var callup model.Callup
	if err := model.DB.Where("id = ? and lord_id = ?", service.ID, user.ID).First(&callup).Error; err != nil {
		c.JSON(http.StatusOK, serializer.Err(serializer.CodeDBError, "召集令查询失败", err))
	}

	if callup.PicturePath == "" {
		c.JSON(http.StatusOK, serializer.Success("无图片文件"))
	} else {
		c.File(callup.PicturePath)
	}
}
