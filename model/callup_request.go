package model

import "github.com/jinzhu/gorm"

// CallupRequest 召集令请求模型
type CallupRequest struct {
	gorm.Model
	CallupID    uint
	Callup      Callup
	RequesterID uint
	Requester   User
	Description string `gorm:"type:text"`
	Status      uint8
}

// 状态码
const (
	Unprocessed uint = iota + 1
	Agreed
	Denied
	Abolished
)
