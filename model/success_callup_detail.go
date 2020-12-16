package model

import "time"

// SuccessCallupDetail 召集令成功明细模型
type SuccessCallupDetail struct {
	CallupID          uint `gorm:"primaryKey;autoIncrement:false"`
	Callup            Callup
	Date              time.Time `gorm:"type:date"`
	SponsorProfit     uint
	ParticipantProfit uint
}
