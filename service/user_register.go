package service

import (
	"call-up/model"
	"call-up/serializer"
)

// UserRegister 用户注册服务
type UserRegister struct {
	UserID       string `form:"user" json:"user" binding:"required,min=4,max=16"`
	Password     string `form:"passwd" json:"passwd" binding:"required,min=6,max=16"`
	Name         string `form:"name" json:"name" binding:"required"`
	IDType       uint8  `form:"id_type" json:"id_type" binding:"required"`
	IDNumber     string `form:"id_number" json:"id_number" binding:"required"`
	Phone        string `form:"phone" json:"phone" binding:"required,len=11"`
	Introduction string `form:"intro" json:"intro"`
	Province     string `form:"province" json:"province" binding:"required"`
	City         string `form:"city" json:"city" binding:"required"`
}

// Register 用户注册
func (service *UserRegister) Register() serializer.Response {
	var count int64 = 0
	if err := model.DB.Model(&model.User{}).Where("user_id = ?", service.UserID).Count(&count).Error; err != nil {
		return serializer.Err(serializer.CodeDBError, "查询用户名失败", err)
	}
	if count > 0 {
		return serializer.Err(serializer.CodeParamErr, "该用户名已经注册", nil)
	}

	user := model.User{
		UserID:       service.UserID,
		Name:         service.Name,
		IDType:       service.IDType,
		IDNumber:     service.IDNumber,
		Phone:        service.Phone,
		Introduction: service.Introduction,
		Province:     service.Province,
		City:         service.City,
	}

	// 加密密码
	if err := user.SetPassword(service.Password); err != nil {
		return serializer.Err(serializer.CodeEncryptError, "密码加密失败", err)
	}

	// 创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		return serializer.Err(serializer.CodeDBError, "注册失败", err)
	}

	resp := serializer.Success("注册成功")
	return resp
}
