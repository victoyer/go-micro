// @Time    : 2019-06-14 20:34
// @Author  : victor！！
// @FileName: token_service.go
// @Software: IntelliJ IDEA

package main

import (
	"github.com/dgrijalva/jwt-go"
	userPKG "shippy-micro/user-service/proto/user"
	"time"
)

type Authable interface {
	Decode(tokenStr string) (*CustomClaims, error)
	Encode(user *userPKG.User) (string, error)
}

// 定义加盐哈希密码所需要的盐，要保证足够的安全可以使用md5
var privateKey = []byte(`xs#a_1-!`)

// 自定义metadata，作为加密后的 JWT 的第二部分返回给客户端
type CustomClaims struct {
	User *userPKG.User
	// 使用标准的 payload
	jwt.StandardClaims
}

type TokenService struct {
	repo Repository
}

func (this *TokenService) Decode(tokenStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return privateKey, nil
	})
	if err != nil {
		return nil, err
	}
	// 解密转换类型返回
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func (this *TokenService) Encode(user *userPKG.User) (string, error) {
	// 设置jwt三天后过期
	expireTime := time.Now().Add(time.Hour * 24 * 3).Unix()
	claims := CustomClaims{
		user,
		jwt.StandardClaims{
			Issuer:    "go.micro.srv.user",
			ExpiresAt: expireTime,
		},
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return jwtToken.SignedString(privateKey)
}
