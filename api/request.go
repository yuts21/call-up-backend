package api

import (
	"call-up/service"
	"github.com/gin-gonic/gin"
	"net/http"
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