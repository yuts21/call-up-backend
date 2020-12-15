package model

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB 数据库链接单例
var DB *gorm.DB

// Database 在中间件中初始化mysql链接
func Database(connString string) {
	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	// Error
	if err != nil {
		log.Panic("连接数据库不成功", err)
	}

	//设置连接池
	sqlDB, err := db.DB()
	if err != nil {
		log.Panic("设置数据库连接池不成功", err)
	}
	//空闲
	sqlDB.SetMaxIdleConns(50)
	//打开
	sqlDB.SetMaxOpenConns(100)
	//超时
	sqlDB.SetConnMaxLifetime(time.Second * 30)

	DB = db

	migration()
}
