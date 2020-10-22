package model

import (
	"github.com/dgrijalva/jwt-go"
)

type Register struct {
	Phone    string `json:"phone" binding:"required"`
	Code     string `json:"code" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// 密码登录
type LoginByP struct {
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// 验证码登录
type LoginByC struct {
	Phone string `json:"phone" binding:"required"`
	Code  string `json:"code" binding:"required"`
}

// 登录返回结果
type LoginResp struct {
	Token     string `json:"token" binding:"required"`
	ExpiresAt int64  `json:"expiresAt" binding:"required"`
}

type Phone struct {
	Phone string `json:"phone" binding:"required"`
}

type Code struct {
	Code string `json:"code" binding:"required"`
}

// jwt
type CustomClaims struct {
	//UUID        uuid.UUID
	//ID          uint
	Username string
	//NickName    string
	//AuthorityId string
	//BufferTime  int64
	jwt.StandardClaims
}
