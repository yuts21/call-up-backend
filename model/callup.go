package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Callup 召集令模型
type Callup struct {
	gorm.Model
	LordID      uint
	Lord        User
	Type        uint8
	Name        string
	Description string `gorm:"type:text"`
	Capacity    uint
	EndDate     time.Time `gorm:"type:date"`
	PicturePath string
	Status      uint8
}

// 状态码
const (
	Waiting uint8 = iota + 1
	Completed
	Expired
	Canceled
)
