package model

//执行数据迁移
func migration() {
	// 自动迁移模式
	DB.AutoMigrate(&User{})
	// 外键约束
	//DB.Model(&Switch{}).AddForeignKey("room_id", "rooms(room_id)", "CASCADE", "CASCADE")
}
