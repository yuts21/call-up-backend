package model

import "log"

const (
	adminID       = "admin"
	adminPassword = "admin"
)

//执行数据迁移
func migration() {
	// 自动迁移模式
	if err := DB.AutoMigrate(&User{}); err != nil {
		log.Fatal("数据模型同步失败", err.Error())
	}
	var count int64 = 0
	DB.Model(&User{}).Where("user_id = ?", adminID).Count(&count)
	if count == 0 {
		var admin = User{
			UserID:   adminID,
			Type:     true,
			Name:     "Administrator",
			IDType:   0,
			Phone:    "00000000000",
			Level:    2,
			Province: "北京市",
			City:     "海淀区",
		}
		if err := admin.SetPassword(adminPassword); err != nil {
			log.Fatal("管理员密码加密失败", err.Error())
		}
		// 创建用户
		if err := DB.Create(&admin).Error; err != nil {
			log.Fatal("管理员添加失败", err.Error())
		}
	}
	if err := DB.AutoMigrate(&Callup{}); err != nil {
		log.Fatal("数据模型同步失败", err.Error())
	}
	if err := DB.AutoMigrate(&Request{}); err != nil {
		log.Fatal("数据模型同步失败", err.Error())
	}
	if err := DB.AutoMigrate(&SuccessCallupDetail{}); err != nil {
		log.Fatal("数据模型同步失败", err.Error())
	}
	if err := DB.AutoMigrate(&AgencyProfit{}); err != nil {
		log.Fatal("数据模型同步失败", err.Error())
	}
}
