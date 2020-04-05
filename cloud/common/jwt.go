package common

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GenerateToken(sub, name string, admin bool) (tokenString string, err error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["sub"] = sub
	claims["name"] = name
	claims["admin"] = admin
	//三个小时后过期
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(3)).Unix()
	token.Claims = claims
	tokenString, err = token.SignedString([]byte(Secret))
	return
}
