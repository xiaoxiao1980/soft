package ini

import (
	"app/global"
	"fmt"
	"os"
)

func Path() {
	// 创建相关目录
	err := os.MkdirAll(global.Config.Path.File, os.ModePerm) //0777
	if err != nil {
		fmt.Println(err)
	}
}
