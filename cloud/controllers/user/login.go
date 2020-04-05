package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/mzz2017/grip/cloud/common"
	"github.com/mzz2017/grip/cloud/common/crypto"
	user2 "github.com/mzz2017/grip/cloud/models/user"
	"github.com/pkg/errors"
)

func PostLogin(ctx *gin.Context) {
	var user user2.User
	ctx.ShouldBindWith(&user, binding.JSON)
	user.Password = crypto.CryptoPwd(user.Password)
	fmt.Print(user)
	if user.Count() == 0 {
		common.ResponseError(ctx, errors.New("用户名或密码错误"))
		return
	}
	//用现有信息去检索完整信息
	user, err := user.Find()
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	tokenString, err := common.GenerateToken(user.Sub, user.Name, user.Admin)
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	common.Response(ctx, common.SUCCESS, gin.H{
		"token": tokenString,
	})
}
