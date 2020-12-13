package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// AgencyProfit 中介收益汇总模型
type AgencyProfit struct {
	gorm.Model
	SuccessDate  time.Time `gorm:"type:date"`
	Province     string
	City         string
	Type         uint8
	CompletedNum uint
	Profit       uint
}
