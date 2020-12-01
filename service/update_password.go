package service

import (
	"call-up/model"
	"call-up/serializer"
	"github.com/gin-gonic/gin"
)

// UpdatePassword 修改密码服务
type UpdatePassword struct {
	Password    string `form:"passwd" json:"passwd" binding:"required,min=6,max=16"`
	NewPassword string `form:"new_passwd" json:"new_passwd" binding:"required,min=6,max=16"`
}

// Update 修改密码函数
func (service *UpdatePassword) Update(c *gin.Context) serializer.Response {
	curUser, _ := c.Get("user")
	user := curUser.(*model.User)

	// 检查旧密码
	if !user.CheckPassword(service.Password) {
		return serializer.Err(serializer.CodeParamErr, "密码错误", nil)
	}

	// 加密新密码
	if err := user.SetPassword(service.NewPassword); err != nil {
		return serializer.Err(serializer.CodeEncryptError, "密码加密失败", err)
	}

	// 修改数据库
	if err := model.DB.Model(&user).Update("password_digest", user.PasswordDigest).Error; err != nil {
		return serializer.Err(serializer.CodeDBError, "修改密码失败", err)
	}

	return serializer.Success("修改密码成功")
}
