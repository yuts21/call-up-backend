package model

import (
	"time"
)

// AgencyProfit 中介收益汇总模型
type AgencyProfit struct {
	ID           uint      `gorm:"primaryKey"`
	SuccessDate  time.Time `gorm:"type:date"`
	Province     string
	City         string
	Type         uint8
	CompletedNum uint
	Profit       uint
}
