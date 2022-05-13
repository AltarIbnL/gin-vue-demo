package common

import (
	"github.com/dgrijalva/jwt-go"
	"myproject/Model"
	"time"
)

// 用于生成加密的密钥

var jwtKey = []byte("a_secret_crect")

type Claim struct {
	Userid uint
	jwt.StandardClaims
}

func ReleaseToken(user Model.User) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour) //TOKEN的有效时间
	claims := &Claim{
		Userid: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "occean.tech",
			Subject:   "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) //选择加密方式
	tokenString, err := token.SignedString(jwtKey)             //使用这个密钥生成token
	if err != nil {
		return "", err
	}
	return tokenString, nil

}

func ParseToken(tokenstr string) (*jwt.Token, *Claim, error) {
	claims := &Claim{}
	token, err := jwt.ParseWithClaims(tokenstr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	//从tokenstr中解析出我们需要的claims然后返回
	return token, claims, err

}
