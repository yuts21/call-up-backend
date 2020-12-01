package service

import (
	"call-up/model"
)

// UserLogin 用户登录服务
type UserLogin struct {
	UserID string `form:"user" json:"user" binding:"required,min=4,max=16"`
	Password string `form:"passwd" json:"passwd" binding:"required,min=5,max=16"`
}

// Login 用户登录函数
func (service *UserLogin) Login() (*model.User, bool) {
	var user model.User
	if err := model.DB.Where("user_id = ?", service.UserID).First(&user).Error; err != nil {
		return nil, false
	}

	if ! user.CheckPassword(service.Password) {
		return nil, false
	}

	return &user, true
}
