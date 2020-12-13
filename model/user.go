package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User 用户模型
type User struct {
	gorm.Model
	UserID         string `gorm:"unique"`
	PasswordDigest string
	Type           bool
	Name           string
	IDType         uint8
	IDNumber       string
	Phone          string `gorm:"size:11"`
	Level          uint8
	Introduction   string `gorm:"type:text"`
	Province       string
	City           string
}

// PassWordCost 密码加密难度
const PassWordCost = 12

// GetUser 用ID获取管理员
func GetUser(ID interface{}) (User, error) {
	var user User
	result := DB.First(&user, ID)
	return user, result.Error
}

// SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}
