package conf

import (
	"call-up/cache"
	"call-up/model"
	"call-up/util"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	godotenv.Load()

	//设置GIN模式
	gin.SetMode(os.Getenv("GIN_MODE"))

	// 设置日志级别
	util.BuildLogger(os.Getenv("LOG_LEVEL"))

	// 读取翻译文件
	if err := LoadLocales("conf/locales/zh-cn.yaml"); err != nil {
		util.Log().Panic("翻译文件加载失败", err)
	}

	// 连接数据库
	model.Database(os.Getenv("MYSQL_DSN"))
	cache.Redis()
}
