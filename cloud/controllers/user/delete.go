package user

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/mzz2017/grip/cloud/common"
	"github.com/mzz2017/grip/cloud/models/user"
	"strconv"
)

func DeleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	err = user.User{Model: gorm.Model{ID: uint(id)}}.Delete()
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	common.Response(ctx, common.SUCCESS, nil)
}
