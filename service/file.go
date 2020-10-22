package service

import (
	. "app/global"
	"app/model"
	"app/pkg/util"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"mime/multipart"

	"go.mongodb.org/mongo-driver/bson"
)

// add 添加某文件类型的路径，默认为all
func AddFilePath(json model.Filepath) error {
	InsertOneResult, err := CollConfig.InsertOne(context.TODO(), bson.D{{"path", json.Path}})
	if err != nil {
		return errors.New("写入缓存失败")
	}
	fmt.Printf("inserted ID: %v\n", InsertOneResult)
	return nil
}

// 保存文件信息到数据库
func SaveFileInfo(file *multipart.FileHeader, fileType string) error {
	// 将文件信息插入 文件集合
	fmt.Println(file.Filename, file.Header, file.Size)
	f, _ := ioutil.ReadFile("./assets/" + file.Filename)
	fmt.Println("string(f)")
	md5 := util.MD5(f)
	// 求file的md5值
	_, err := CollFile.InsertOne(context.TODO(), bson.D{
		{"name", file.Filename},
		{"size", file.Size},
		{"md5", md5},
		{"fileType", fileType},
		{"path", "./assets/" + file.Filename}})
	if err != nil {
		return err
	}
	return nil
}
