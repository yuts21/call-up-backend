package model

import (
	"call-up/cache"
	"strconv"
	"time"

	"gorm.io/gorm"
)

// Callup 召集令模型
type Callup struct {
	gorm.Model
	SponsorID   uint
	Sponsor     User `gorm:"foreignKey:SponsorID"`
	Type        uint8
	Name        string
	Description string `gorm:"type:text"`
	Capacity    uint
	EndDate     time.Time `gorm:"type:date"`
	PicturePath string
	Canceled    bool
	Request     []Request
}

// 状态码
const (
	Waiting uint8 = iota + 1
	Completed
	Expired
	Canceled
)

// Status 召集令状态
func (callup *Callup) Status() uint8 {
	if callup.Canceled {
		return Canceled
	}

	var count int64 = 0
	strCallupID := "callup_" + strconv.FormatUint(uint64(callup.ID), 10)
	if cache.RedisClient.Exists(strCallupID).Val() == 0 {
		DB.Model(&Request{}).Where("callup_id = ? and status = ?", callup.ID, Agreed).Count(&count)
		cache.RedisClient.Set(strCallupID, strconv.FormatInt(count, 10), 0)
	} else {
		count, _ = strconv.ParseInt(cache.RedisClient.Get(strCallupID).Val(), 10, 64)
	}
	if uint(count) >= callup.Capacity {
		return Completed
	}

	year, month, day := time.Now().Date()
	curDay := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	if curDay.Before(callup.EndDate) || curDay.Equal(callup.EndDate) {
		return Waiting
	}
	return Expired
}
