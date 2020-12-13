package api

import (
	"call-up/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AgencyProfitList 请求信息列表
func AgencyProfitList(c *gin.Context) {
	var serv service.AgencyProfitList
	if err := c.ShouldBind(&serv); err == nil {
		res := serv.List(c)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}
