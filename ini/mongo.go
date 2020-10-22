package ini

import (
	"app/global"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// 初始化数据库
func Mongo() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://@localhost:27017"))
	if err != nil {
		log.Print(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Print(err)
	}

	// 设置全局数据库
	name := "hunan"
	global.Gmg = client.Database(name)

	// todo 删除
	global.CollUser = client.Database(name).Collection("user")
	global.CollFile = client.Database(name).Collection("file")
	global.CollCache = client.Database(name).Collection("cache")
	global.CollConfig = client.Database(name).Collection("config")

	fmt.Println("mongo")
}
