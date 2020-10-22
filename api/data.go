package api

import (
	. "app/global"
	"app/global/resp"
	"app/model"
	"log"

	"github.com/gin-gonic/gin"
)

// @Summary 档案字段
// @Description .
// @Tags 读取数据
// @Success 200 {string} string "{"message": "ok"}"
// @Router /data/field [get]
// @Security Bearer
func GetField(c *gin.Context) {
	var archive []model.Field
	result := Gdb.Find(&archive)
	log.Println("Field", result)
	resp.OkWithData(archive, c)
}

// @Summary 档案文件信息
// @Description 。
// @Tags 读取数据
// @Success 200 {string} string "{"message": "ok"}"
// @Router /data/file [get]
// @Security Bearer
func GetFile(c *gin.Context) {
	var archive []model.File
	result := Gdb.Find(&archive)
	log.Println("File", result)
	resp.OkWithData(archive, c)
}

// @Summary 档案版本信息
// @Description 。
// @Tags 读取数据
// @Success 200 {string} string "{"message": "ok"}"
// @Router /data/version [get]
// @Security Bearer
func GetVersion(c *gin.Context) {
	var archive []model.Version
	result := Gdb.Find(&archive)
	log.Println("Version", result)
	resp.OkWithData(archive, c)
}

// @Summary 档案记录
// @Description 。
// @Tags 读取数据
// @Success 200 {string} string "{"message": "ok"}"
// @Router /data/record [get]
// @Security Bearer
func GetRecord(c *gin.Context) {
	var archive []model.Record
	result := Gdb.Find(&archive)
	log.Println("Record", result)
	resp.OkWithData(archive, c)
}

// @Summary 获取表的全部记录
// @Description 。
// @Tags 读取数据
// @Success 200 {string} string "{"message": "ok"}"
// @Router /data/table [get]
// @Security Bearer
func GetTable(c *gin.Context) {
	var archive []model.Result
	Gdb.Raw("SELECT * FROM archivemanage_data.tt_archive_field").Scan(&archive)
	log.Println("GetTable", archive)
	resp.OkWithData(archive, c)
}
