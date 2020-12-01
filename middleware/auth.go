package middleware

import (
	"call-up/model"
	"call-up/serializer"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserAuth 普通用户鉴权
func UserAuth() gin.HandlerFunc {
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

// AdminAuth 管理员用户鉴权
func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		curUser, _ := c.Get("user")
		if curUser.(*model.User).Type {
			c.Next()
			return
		}
		c.JSON(http.StatusUnauthorized, serializer.Err(serializer.CodeCheckLogin, "非管理员用户", nil))
		c.Abort()
	}
}
