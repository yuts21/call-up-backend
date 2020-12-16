package api

import (
	"call-up/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AgencyProfitList 中介收益列表
func AgencyProfitList(c *gin.Context) {
	var serv service.AgencyProfitList
	if err := c.ShouldBind(&serv); err == nil {
		res := serv.List(c)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

// CallupProfitList 召集令收益列表
func CallupProfitList(c *gin.Context) {
	var serv service.CallupProfitList
	if err := c.ShouldBind(&serv); err == nil {
		res := serv.List(c)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}
