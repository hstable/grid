package tree

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mzz2017/grip/cloud/common"
	"github.com/mzz2017/grip/cloud/models/tower"
	"strconv"
)

func GetTowers(ctx *gin.Context) {
	_id := ctx.Param("id")
	id, err := strconv.Atoi(_id)
	if _id == "" || err != nil {
		common.ResponseError(ctx, errors.New("invalid id"))
		return
	}
	t, err := tower.GetByLineID(id)
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	common.ResponseSuccess(ctx, gin.H{
		"towers": t,
	})
}
