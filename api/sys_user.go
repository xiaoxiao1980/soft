package api

import (
	"app/global"
	"app/global/resp"
	"app/model"
	"app/pkg/middleware"
	"app/pkg/util"
	"context"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type RegisterAndLoginStruct struct {
	Username  string `json:"username" bson:"username"`
	Password  string `json:"password" bson:"ps"`
	Captcha   string `json:"captcha"`
	CaptchaId string `json:"captchaId"`
}

// @Summary 管理用户注册
// @Description 用户注册，填写用户呢称和密码
// @Tags system
// @Success 200 {string} string "{"message": "ok"}"
// @Router /sys/register [get]
// @Security Bearer
func SysRegister(c *gin.Context) {
	var json RegisterAndLoginStruct
	_ = c.ShouldBindJSON(&json)
	json.Password = util.MD5V(json.Password)
	global.Gmg.Collection("sys_user").InsertOne(context.TODO(), json)
	resp.Ok(c)
}

// @Summary 管理用户登录
// @Description 用户登录，填写用户呢称和密码
// @Tags system
// @Success 200 {string} string "{"message": "ok"}"
// @Router /sys/login [get]
// @Security Bearer
func SysLogin(c *gin.Context) {
	var json RegisterAndLoginStruct
	_ = c.ShouldBindJSON(&json)

	var result bson.M
	err := global.Gmg.Collection("sys_user").FindOne(context.TODO(), bson.D{{"username", json.Username}, {"ps", util.MD5V(json.Password)}}).Decode(&result)
	if err != nil {
		fmt.Println(err)
		resp.OkDetailed(err, "账号不存在或密码错误", c)
	} else {
		tokenSysNext(c, &json)
	}
}

//登录以后签发jwt
func tokenSysNext(c *gin.Context, json *RegisterAndLoginStruct) {
	j := &middleware.Jwt{
		SigningKey: []byte(global.Config.JWT.SigningKey), // 不加上SigningKey会提示 composite literal uses unkeyed fields
	}
	//j := middleware.NewJWT()
	clams := model.CustomClaims{
		Username: json.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 60*60*24*7, // 过期时间 一周
			Issuer:    "tdkj",                         // 签名的发行者
			NotBefore: time.Now().Unix() - 1000,       // 签名生效时间
		},
	}
	token, err := j.CreateToken(clams)
	if err != nil {
		resp.FailWithMessage("获取token失败", c)
	} else {
		resp.OkWithData(model.LoginResp{
			Token:     token,
			ExpiresAt: clams.StandardClaims.ExpiresAt * 1000,
		}, c)
	}
}
