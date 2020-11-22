package main

import (
	"call-up/conf"
	"call-up/server"
	"log"
)

func main() {
	// 从配置文件读取配置
	conf.Init()

	// 装载路由
	r := server.NewRouter()
	if err := r.Run(":3000"); err != nil {
		log.Fatal(err.Error())
	}
}
