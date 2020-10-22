package main

import (
	"app/ini"
)

func main() {
	// 连接mongodb
	ini.Mongo()

	// 连接mysql
	ini.GormMysql()

	// 检查保存路径，若不存在则创建
	// ini.Path()

	// 创建gin路由 ,最后初始化
	ini.RunWebServer()
}
