package user

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mzz2017/grip/cloud/common"
	"github.com/mzz2017/grip/cloud/models/user"
)

func PostUser(ctx *gin.Context) {
	var newUser user.User
	err := ctx.BindJSON(&newUser)
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	// 数据库是否存在相同的Sub
	if u := (user.User{Sub: newUser.Sub}); u.Count() > 0 {
		common.ResponseError(ctx, errors.New("用户名已存在"))
		return
	}
	err = newUser.Insert()
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	// 生成身份token返给前端
	tokenString, err := common.GenerateToken(newUser.Sub, newUser.Name, newUser.Admin)
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	common.Response(ctx, common.SUCCESS, gin.H{
		"token": tokenString,
	})
}
