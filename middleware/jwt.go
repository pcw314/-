package service

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// securityKey 随机UUID当密钥
var securityKey = []byte("EFBB4BE4-E8BB-6C04-F090-3A86291E7D10")

type UserClaims struct {
	UserId   int    `json:"user_id"`
	Role     int    `json:"role"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// MakeToken 生成Token
func MakeToken(userId int, username string, role int) (token string, err error) {
	claim := UserClaims{
		UserId:   userId,
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			// 签发者
			Issuer: "WiseCloud",
			// 签发对象
			Subject: "NiuB",
			// 过期时间三小时
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)), //一周后过期
			// 签发时间
			IssuedAt: jwt.NewNumericDate(time.Now()),
			// 生效时间
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	// 使用HS256算法
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	token, err = t.SignedString(securityKey)
	if err != nil {
		return "", err
	}
	return
}

// ParseToken 解析Token
func ParseToken(token string) (*UserClaims, error) {
	t, err := jwt.ParseWithClaims(token, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return securityKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !t.Valid {
		return nil, errors.New("claim invalid")
	}
	claims, ok := t.Claims.(*UserClaims)
	if !ok {
		return nil, errors.New("invalid claim type")
	}
	return claims, nil
}
