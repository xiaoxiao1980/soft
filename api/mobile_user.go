package api

import (
	"app/global"
	"app/global/resp"
	"app/model"
	"app/pkg/middleware"
	"app/pkg/util"
	"app/service"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// @Summary 发送验证码
// @Description  直接发送验证码给前端，以后的话用第三方发送给用户手机
// @Tags mobile
// @Success 200 {string} string "{"message": "ok"}"
// @Router /base/code [get]
// @Security Bearer
func SendCode(c *gin.Context) {
	var json model.Phone
	if err := c.ShouldBindJSON(&json); err != nil {
		resp.Fail(c)
		return
	}
	fmt.Println("R", json, json.Phone)
	// 生成六位验证码
	num := util.CreateCaptcha()
	err := service.SaveCode(json.Phone, num)
	if err != nil {
		resp.FailWithMessage(fmt.Sprintf("%v", err), c)
	} else {
		resp.OkWithData(model.Code{Code: num}, c)
	}
}

// @Summary 手机用户注册
// @Description 手机号密码和验证码
// @Tags mobile
// @Success 200 {string} string "{"message": "ok"}"
// @Router /mobile/register [post]
// @Security Bearer
func Register(c *gin.Context) {
	var json model.Register
	if err := c.ShouldBindJSON(&json); err != nil {
		resp.Fail(c)
		return
	}
	result, err := service.Register(json)
	fmt.Println(result)
	if err != nil {
		resp.FailWithDetailed(resp.ERROR, result, fmt.Sprintf("%v", err), c)
	} else {
		resp.OkDetailed(*result, "注册成功", c)
	}
}

// @Summary 登录
// @Description 手机用户登录
// @Tags mobile
// @Success 200 {string} string "{"message": "ok"}"
// @Router /mobile/login [get]
// @Security Bearer
func Login(c *gin.Context) {
	var json model.LoginByP
	if err := c.ShouldBindJSON(&json); err != nil {
		resp.Fail(c)
		return
	}

	if err := service.Login(json); err != nil {
		resp.FailWithMessage(fmt.Sprintf("%v", err), c)
	} else {
		// resp.Ok(c)
		tokenNext(c, &json)
	}
}

//登录以后签发jwt
func tokenNext(c *gin.Context, json *model.LoginByP) {
	j := &middleware.Jwt{
		SigningKey: []byte(global.Config.JWT.SigningKey), // 不加上SigningKey会提示 composite literal uses unkeyed fields
	}
	clams := model.CustomClaims{
		// UUID: user.UUID, // 分布式用到
		// ID:   user.ID,
		// NickName:    user.NickName,
		// AuthorityId: user.AuthorityId,
		StandardClaims: jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000),       // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 60*60*24*7), // 过期时间 一周
			Issuer:    "tdkj",                                // 签名的发行者
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
