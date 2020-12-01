package model

import "github.com/jinzhu/gorm"

// Request 召集令请求模型
type Request struct {
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
	RequestUnprocessed uint8 = iota
	RequestAgreed
	RequestDenied
	RequestCanceled
)
