package api

import (
	"call-up/service"
	"github.com/gin-gonic/gin"
)

// CallupCreate 创建召集令
func CallupCreate(c *gin.Context) {
	var serv service.CallupCreate
	if err := c.ShouldBind(&serv); err == nil {
		res := serv.Create(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
