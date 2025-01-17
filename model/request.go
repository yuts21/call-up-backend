package model

import "gorm.io/gorm"

// Request 接令请求模型
type Request struct {
	gorm.Model
	CallupID    uint
	Callup      Callup
	RequesterID uint
	Requester   User   `gorm:"foreignKey:RequesterID"`
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
