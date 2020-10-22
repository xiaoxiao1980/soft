package api

import (
	"app/global/resp"
	"app/service"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

// 上传文件并保存到目录
//func Upload(c *gin.Context)  {
//	file, _ := c.FormFile("hello")
//	fmt.Println(file.Filename)
//	fmt.Println(file.Header)
//	fmt.Println(file.Size)
//
//	dst := fmt.Sprintf(savePath+file.Filename)
//	err :=c.SaveUploadedFile(file, dst)
//	if err!=nil{
//		fmt.Println(err)
//	}
//	resp.Ok(c)
//}

type FileType struct {
	Type string `uri:"type" binding:"required"`
}

// @Summary 监控
// @Description 上传多个文件
// @Tags 文件
// @Success 200 {string} string "{"message": "ok"}"
// @Router /file/upload/type [post]
// @Security Bearer
func MultiUpload(c *gin.Context) {
	var json FileType
	if err := c.ShouldBindUri(&json); err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"msg": err})
		return
	}
	form, _ := c.MultipartForm()
	files := form.File["upload[]"]

	for _, file := range files {
		// 保存文件到服务器（assets）
		if err := service.SaveFileInfo(file, json.Type); err != nil {
			fmt.Println(err)
		}
		// 相对路径
		dst := fmt.Sprintf("./assets/" + file.Filename)
		fmt.Println("目标保存位置", dst)
		c.SaveUploadedFile(file, dst)
	}

	resp.Ok(c)
}

// 获取所有文件
func GetAllFile(c *gin.Context) {

}
