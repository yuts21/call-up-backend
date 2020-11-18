package service

import (
	"call-up/model"
	"call-up/serializer"
)

// UserRegisterService 用户注册服务
type UserRegisterService struct {
	UserID string `form:"user_id" json:"user_id" binding:"required,min=4,max=16"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=16"`
}

// valid 验证表单
func (service *UserRegisterService) valid() *serializer.Response {
	count := 0
	model.DB.Model(&model.User{}).Where("user_id = ?", service.UserID).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Code: serializer.CodeParamErr,
			Msg:  "该用户名已经注册",
		}
	}
	return nil
}

// Register 用户注册
func (service *UserRegisterService) Register() serializer.Response {
	user := model.User{
		UserID: service.UserID,
	}

	// 表单验证
	if err := service.valid(); err != nil {
		return *err
	}

	// 加密密码
	if err := user.SetPassword(service.Password); err != nil {
		return serializer.Err(serializer.CodeEncryptError, "密码加密失败", err)
	}

	// 创建管理员
	if err := model.DB.Create(&user).Error; err != nil {
		return serializer.Err(serializer.CodeDBError, "注册失败", err)
	}

	resp := serializer.BuildUserResponse(user)
	resp.Msg = "注册成功"
	return resp
}
