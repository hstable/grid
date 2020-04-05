package user

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/mzz2017/grip/cloud/common"
	"github.com/mzz2017/grip/cloud/models/user"
)

func PostRegister(ctx *gin.Context) {
	params := struct {
		Sub      string `binding:"required"`
		Name     string `binding:"required"`
		Admin    bool   `binding:"required"`
		Password string `binding:"required"`
	}{}
	err := ctx.ShouldBindWith(&params, binding.JSON)
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	// 数据库是否存在相同的Sub
	if u := (user.User{Sub: params.Sub}); u.Count() > 0 {
		common.ResponseError(ctx, errors.New("用户名已存在"))
		return
	}
	// 在数据库插入数据，密码在Insert函数里面已经加密好
	newUser := user.User{
		Sub:      params.Sub,
		Admin:    params.Admin,
		Name:     params.Name,
		Password: params.Password,
	}
	err = newUser.Insert()
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	// 生成身份token返给前端
	tokenString, err := common.GenerateToken(params.Sub, params.Name, params.Admin)
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	common.Response(ctx, common.SUCCESS, gin.H{
		"token": tokenString,
	})
}
