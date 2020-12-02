package model

import "github.com/jinzhu/gorm"

// Request 接令请求模型
type Request struct {
	gorm.Model
	Callup      Callup
	CallupID    uint
	Requester   User
	RequesterID uint
	Description string `gorm:"type:text"`
	Status      uint8
}

// 状态码
const (
	Unprocessed uint8 = iota + 1
	Agreed
	Denied
	Abolished
)
