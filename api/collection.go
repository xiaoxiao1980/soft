package api

import (
	"app/global/resp"
	"fmt"

	"github.com/gin-gonic/gin"
)

func InsertOne(c *gin.Context) {
	fmt.Println("插入一条数据")
	resp.Ok(c)
}

func InsertMany(c *gin.Context) {
	resp.Ok(c)
}

func Find(c *gin.Context) {
	resp.Ok(c)
}

func FindOne(c *gin.Context) {
	resp.Ok(c)
}

func FindOneAndDelete(c *gin.Context) {
	resp.Ok(c)
}
func FindOneAndReplace(c *gin.Context) {
	resp.Ok(c)
}

func FindOneAndUpdate(c *gin.Context) {
	resp.Ok(c)
}

func UpdateOne(c *gin.Context) {
	resp.Ok(c)
}
func UpdateMany(c *gin.Context) {
	resp.Ok(c)
}

func DeleteOne(c *gin.Context) {
	resp.Ok(c)
}

func DeleteMany(c *gin.Context) {
	resp.Ok(c)
}

func Aggregate(c *gin.Context) {
	resp.Ok(c)
}

func BulkWrite(c *gin.Context) {
	resp.Ok(c)
}

func CountDocuments(c *gin.Context) {
	resp.Ok(c)
}

func Distinct(c *gin.Context) {
	resp.Ok(c)
}

func EstimatedDocumentCount(c *gin.Context) {
	resp.Ok(c)
}

func ReplaceOne(c *gin.Context) {
	resp.Ok(c)
}

func Watch(c *gin.Context) {
	resp.Ok(c)
}
