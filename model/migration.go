package model

import "log"

const (
	adminID       = "admin"
	adminPassword = "admin"
)

//执行数据迁移
func migration() {
	// 自动迁移模式
	DB.AutoMigrate(&User{})
	count := 0
	DB.Model(&User{}).Where("user_id = ?", adminID).Count(&count)
	if count == 0 {
		var admin = User{
			UserID:  adminID,
			Type:    true,
			Name:    "Administrator",
			IDType:  0,
			Phone:   "00000000000",
			Level:   2,
			RegCity: "北京市海淀区",
		}
		if err := admin.SetPassword(adminPassword); err != nil {
			log.Fatal("管理员密码加密失败", err)
		}
		// 创建用户
		if err := DB.Create(&admin).Error; err != nil {
			log.Fatal("管理员添加失败", err)
		}
	}
	DB.AutoMigrate(&Callup{})
	// 外键约束
	//DB.Model(&Switch{}).AddForeignKey("room_id", "rooms(room_id)", "CASCADE", "CASCADE")
}
