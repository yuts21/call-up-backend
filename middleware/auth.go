package middleware

import (
	"call-up/model"
	"call-up/serializer"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PlayerAuth 普通用户鉴权
func PlayerAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		curUser, _ := c.Get("user")
		if !curUser.(*model.User).Type {
			c.Next()
			return
		}
		c.JSON(http.StatusUnauthorized, serializer.Err(serializer.CodeCheckLogin, "非普通用户", nil))
		c.Abort()
	}
}
