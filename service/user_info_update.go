package service

import (
	"call-up/model"
	"call-up/serializer"
	"github.com/gin-gonic/gin"
)

// UserInfoUpdate 修改用户信息服务
type UserInfoUpdate struct {
	Phone        *string `form:"phone" json:"phone" binding:"omitempty,len=11"`
	Introduction *string `form:"intro" json:"intro"`
}

// Update 修改用户信息函数
func (service *UserInfoUpdate) Update(c *gin.Context) serializer.Response {
	curUser, _ := c.Get("user")
	user := curUser.(*model.User)

	// 修改数据库
	userNew := make(map[string]interface{})
	if service.Phone != nil {
		userNew["phone"] = *service.Phone
	}
	if service.Introduction != nil {
		userNew["introduction"] = *service.Introduction
	}
	if err := model.DB.Model(&user).Updates(userNew).Error; err != nil {
		return serializer.Err(serializer.CodeDBError, "修改用户信息失败", err)
	}

	return serializer.Success("修改用户信息成功")
}
