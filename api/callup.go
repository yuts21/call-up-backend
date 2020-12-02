package api

import (
	"call-up/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CallupCreate 创建召集令
func CallupCreate(c *gin.Context) {
	var serv service.CallupCreate
	if err := c.ShouldBind(&serv); err == nil {
		res := serv.Create(c)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

// CallupInfo 召集令信息查询
func CallupInfo(c *gin.Context) {
	var serv service.CallupInfo
	if err := c.ShouldBind(&serv); err == nil {
		res := serv.Info(c)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

// CallupPicture 召集令图片查询
func CallupPicture(c *gin.Context) {
	var serv service.CallupPicture
	if err := c.ShouldBind(&serv); err == nil {
		serv.GetPicture(c)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}