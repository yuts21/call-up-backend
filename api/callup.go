package api

import (
	"call-up/service"
	"net/http"

	"github.com/gin-gonic/gin"
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

// CallupMine 自己发布的召集令列表查询
func CallupMine(c *gin.Context) {
	var serv service.CallupMine
	if err := c.ShouldBind(&serv); err == nil {
		res := serv.List(c)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

// CallupList 召集令列表查询
func CallupList(c *gin.Context) {
	var serv service.CallupList
	if err := c.ShouldBind(&serv); err == nil {
		res := serv.List(c)
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

// CallupCancel 召集令取消
func CallupCancel(c *gin.Context) {
	var serv service.CallupCancel
	if err := c.ShouldBind(&serv); err == nil {
		res := serv.Cancel(c)
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

// CallupRequestInfo 召集令信息列表
func CallupRequestInfo(c *gin.Context) {
	var serv service.CallupRequestInfo
	if err := c.ShouldBind(&serv); err == nil {
		res := serv.Info(c)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

// CallupRequestList 召集令请求列表
func CallupRequestList(c *gin.Context) {
	var serv service.CallupRequestList
	if err := c.ShouldBind(&serv); err == nil {
		res := serv.List(c)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

// CallupRequestHandle 召集令请求处理
func CallupRequestHandle(c *gin.Context) {
	var serv service.CallupRequestHandle
	if err := c.ShouldBind(&serv); err == nil {
		res := serv.Handle(c)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}
