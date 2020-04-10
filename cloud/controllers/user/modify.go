package user

import (
	"github.com/gin-gonic/gin"
	"github.com/mzz2017/grip/cloud/common"
	"github.com/mzz2017/grip/cloud/models/user"
	"strconv"
)

func PutUser(ctx *gin.Context) {
	var newUser user.User
	err := ctx.BindJSON(&newUser)
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	newUser.ID = uint(id)
	err = newUser.Update()
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	common.Response(ctx, common.SUCCESS, nil)
}
