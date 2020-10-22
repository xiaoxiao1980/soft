package service

import (
	. "app/global"
	"app/model"
	"app/pkg/util"
	"context"
	"errors"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// 把手机号和验证码放入缓存中
func SaveCode(phone, num string) error {
	result, err := CollCache.InsertOne(context.TODO(), bson.D{{"phone", phone}, {"code", num}})
	if err != nil {
		return errors.New("写入缓存失败")
	}
	fmt.Printf("inserted ID: %v\n", result.InsertedID)
	return nil
}

func Register(json model.Register) (*mongo.InsertOneResult, error) {
	// 1.判断手机号是否存在
	// 2.如果未存在，则把发送过来的验证码和cache中的对比，一致则创建用户
	fmt.Println("json", json)
	cursor, err := CollUser.Find(context.TODO(), bson.D{{"phone", json.Phone}})
	if err != nil {
		return nil, err
	}
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}
	// 如果已经注册
	if len(results) != 0 {
		return nil, errors.New("该用户已经存在")
	} else {
		// TO-DO 判断缓存中的验证码是否和发送过来的一致，一致则插入
		InsertOneResult, err := CollUser.InsertOne(context.TODO(), bson.D{{"phone", json.Phone}, {"password", util.MD5V(json.Password)}})
		if err != nil {
			return nil, errors.New("写入缓存失败")
		}
		fmt.Printf("inserted ID: %v\n", InsertOneResult)
		return InsertOneResult, nil
	}
}

// 密码登录
func Login(json model.LoginByP) error {
	// 查找数据库，看发送来的是否和数据库里面的一致
	var result bson.M
	err := CollUser.FindOne(context.TODO(), bson.D{{"phone", json.Phone}}).Decode(&result)
	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection 就是没找到
		if err == mongo.ErrNoDocuments {
			return errors.New("手机号未注册")
		}
		log.Fatal(err)
	}
	fmt.Println("result", result["password"])
	// 密码一致，登录成功
	if result["password"] == util.MD5V(json.Password) {
		return nil
	}
	return errors.New("密码错误")
}
