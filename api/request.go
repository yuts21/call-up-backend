package api

import (
	"call-up/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RequestCreate 创建请求
func RequestCreate(c *gin.Context) {
	var serv service.RequestCreate
	if err := c.ShouldBind(&serv); err == nil {
		res := serv.Create(c)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

// RequestInfo 请求信息查询
func RequestInfo(c *gin.Context) {
	var serv service.RequestInfo
	if err := c.ShouldBind(&serv); err == nil {
		res := serv.Info(c)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

// RequestList 请求信息列表
func RequestList(c *gin.Context) {
	var serv service.RequestList
	if err := c.ShouldBind(&serv); err == nil {
		res := serv.List(c)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

// RequestAll 全部请求信息列表
func RequestAll(c *gin.Context) {
	var serv service.RequestAll
	if err := c.ShouldBind(&serv); err == nil {
		res := serv.List(c)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

// RequestUpdate 请求信息修改
func RequestUpdate(c *gin.Context) {
	var serv service.RequestUpdate
	if err := c.ShouldBind(&serv); err == nil {
		res := serv.Update(c)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

// RequestCancel 请求信息取消
func RequestCancel(c *gin.Context) {
	var serv service.RequestCancel
	if err := c.ShouldBind(&serv); err == nil {
		res := serv.Cancel(c)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

// RequestDelete 请求信息删除
func RequestDelete(c *gin.Context) {
	var serv service.RequestDelete
	if err := c.ShouldBind(&serv); err == nil {
		res := serv.Delete(c)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

// RequestCallup 已接令召集令查询
func RequestCallup(c *gin.Context) {
	var serv service.RequestCallup
	if err := c.ShouldBind(&serv); err == nil {
		res := serv.List(c)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}
