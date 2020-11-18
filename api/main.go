package api

import (
	"call-up/conf"
	"call-up/serializer"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Ping 状态检查页面
func Ping(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "Pong",
	})
}

// ErrorResponse 返回错误消息
func ErrorResponse(err error) serializer.Response {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			field := conf.T(fmt.Sprintf("Field.%s", e.Field))
			tag := conf.T(fmt.Sprintf("Tag.Valid.%s", e.Tag))
			return serializer.Err(serializer.CodeParamErr, fmt.Sprintf("%s%s", field, tag), err)
		}
	}
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.Err(serializer.CodeParamErr, "JSON类型不匹配", err)
	}

	return serializer.Err(serializer.CodeParamErr, "参数错误", err)
}
