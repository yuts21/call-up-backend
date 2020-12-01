package conf

import (
	"call-up/cache"
	"call-up/model"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	if err := godotenv.Load(); err != nil {
		log.Fatal("环境变量加载失败", err.Error())
	}

	//设置GIN模式
	gin.SetMode(os.Getenv("GIN_MODE"))

	// 读取翻译文件
	if err := LoadLocales("conf/locales/zh-cn.yaml"); err != nil {
		log.Fatal("翻译文件加载失败", err.Error())
	}

	// 连接数据库
	model.Database(os.Getenv("MYSQL_DSN"))
	cache.Redis()
}
