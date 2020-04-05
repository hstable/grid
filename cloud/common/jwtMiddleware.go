package common

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
	"github.com/mzz2017/grip/cloud/config"
	user2 "github.com/mzz2017/grip/cloud/models/user"
	"github.com/pkg/errors"
)

var Secret string

func init() {
	Secret = config.Get().Secret
}

func JWTAuth(Admin bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := request.ParseFromRequest(ctx.Request, request.AuthorizationHeaderExtractor,
			func(token *jwt.Token) (interface{}, error) {
				// 我们使用固定的secret，直接返回就好
				return []byte(Secret), nil
			})
		if err != nil {
			Response(ctx, UNAUTHORIZED, err.Error())
			ctx.Abort()
			return
		}
		if !token.Valid {
			Response(ctx, UNAUTHORIZED, "Token is invalid!")
			ctx.Abort()
			return
		}
		//如果需要Admin权限
		mapClaims := token.Claims.(jwt.MapClaims)
		if Admin && mapClaims["admin"] == false {
			Response(ctx, UNAUTHORIZED, "Token is invalid!")
			ctx.Abort()
			return
		}
		//将名字和sub丢入参数
		ctx.Set("Name", mapClaims["name"])
		ctx.Set("Sub", mapClaims["sub"])
		user, err := user2.User{Sub: mapClaims["sub"].(string)}.Find()
		if err != nil {
			Response(ctx, UNAUTHORIZED, errors.New("用户不存在"))
			ctx.Abort()
			return
		}
		ctx.Set("User", user)
		//在ctx.Next()前的都是before request，之后的是after request
		ctx.Next()
	}
}
