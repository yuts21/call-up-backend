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

// CallupUpdate 召集令修改
func CallupUpdate(c *gin.Context) {
	var serv service.CallupUpdate
	if err := c.ShouldBind(&serv); err == nil {
		res := serv.Update(c)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

// CallupDelete 召集令删除
func CallupDelete(c *gin.Context) {
	var serv service.CallupDelete
	if err := c.ShouldBind(&serv); err == nil {
		res := serv.Delete(c)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}
