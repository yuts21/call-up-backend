package service

import (
	"call-up/model"
	"call-up/serializer"
	"github.com/gin-gonic/gin"
	"net/http"
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
	if err := model.DB.Where("id = ?", service.ID).First(&callup).Error; err != nil {
		c.JSON(http.StatusOK, serializer.Err(serializer.CodeDBError, "召集令查询失败", err))
	}
	if !user.Type && user.ID != callup.LordID {
		c.JSON(http.StatusOK, serializer.Err(serializer.CodeNoRightErr, "无权限", nil))
	}

	if callup.PicturePath == "" {
		c.JSON(http.StatusOK, serializer.Success("无图片文件"))
	} else {
		c.File(callup.PicturePath)
	}
}
