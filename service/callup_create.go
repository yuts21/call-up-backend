package service

import (
	"call-up/conf"
	"call-up/model"
	"call-up/serializer"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// CallupCreate 创建召集令服务
type CallupCreate struct {
	Type        uint8  `form:"type" json:"type" binding:"required"`
	Name        string `form:"name" json:"name" binding:"required"`
	Description string `form:"descrpt" json:"descrpt"`
	Capacity    uint   `form:"cap" json:"cap" binding:"required,gt=0"`
	EndDate     int64  `form:"end_date" json:"end_date" binding:"required"`
}

// Create 创建召集令
func (service *CallupCreate) Create(c *gin.Context) serializer.Response {
	curUser, _ := c.Get("user")
	Sponsor := curUser.(*model.User)

	callup := model.Callup{
		SponsorID:   Sponsor.ID,
		Type:        service.Type,
		Name:        service.Name,
		Description: service.Description,
		Capacity:    service.Capacity,
		EndDate:     time.Unix(service.EndDate, 0),
	}

	if err := model.DB.Create(&callup).Error; err != nil {
		return serializer.Err(serializer.CodeDBError, "召集令创建失败", err)
	}

	picture, err := c.FormFile("pic")
	if err != nil {
		return serializer.Success("召集令创建成功")
	}

	picName := "callup_" + strconv.FormatUint(uint64(callup.ID), 10) + filepath.Ext(picture.Filename)
	picName = path.Join(conf.FilePath, picName)
	if err := c.SaveUploadedFile(picture, picName); err != nil {
		return serializer.Err(serializer.CodeFileUploadError, "照片上传失败", err)
	}

	if err := model.DB.Model(&callup).Update("picture_path", picName).Error; err != nil {
		_ = os.Remove(picName)
		return serializer.Err(serializer.CodeDBError, "照片上传失败", err)
	}

	return serializer.Success("召集令创建成功")
}
