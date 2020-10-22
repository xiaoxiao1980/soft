package api

import (
	"app/global/resp"
	"app/model"
	"app/service"
	"fmt"
	"os"
	"path"

	"github.com/gin-gonic/gin"
)

func Select(c *gin.Context) {
	var json model.Filepath
	if err := c.ShouldBindJSON(&json); err != nil {
		resp.Fail(c)
		return
	}
	ePath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	// 全路径
	fmt.Println("全路径", json)
	if err = service.AddFilePath(json); err != nil {
		resp.Fail(c)
	}

	// 全路径
	fmt.Println("全路径", ePath)
	// 所在目录
	fmt.Println("当前路径", path.Dir(ePath))
	resp.Ok(c)
}
