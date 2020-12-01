package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

// Callup 召集令模型
type Callup struct {
	gorm.Model
	LordID      uint
	Lord        User
	Type        uint8
	Name        string
	Description string `gorm:"type:text"`
	NumPerson   uint
	EndDate     time.Time `gorm:"type:date"`
	Picture     []byte    `gorm:"type:blob"`
	Status      uint8
}
