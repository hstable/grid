package tree

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mzz2017/grip/cloud/common"
	"github.com/mzz2017/grip/cloud/models/image"
	"strconv"
)

func getParams(ctx *gin.Context) (id, after, limit int, err error) {
	_id := ctx.Param("id")
	_after := ctx.DefaultQuery("after", "0")
	_limit := ctx.DefaultQuery("limit", "4")
	id, err = strconv.Atoi(_id)
	if _id == "" || err != nil {
		err = errors.New("invalid id")
		return
	}
	after, err = strconv.Atoi(_after)
	if _after == "" || err != nil {
		err = errors.New("invalid after")
		return
	}
	limit, err = strconv.Atoi(_limit)
	if _limit == "" || err != nil {
		err = errors.New("invalid limit")
		return
	}
	return
}

func GetImages(ctx *gin.Context) {
	id, after, limit, err := getParams(ctx)
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	t, err := image.GetByTowerIDAfter(id, after, limit)
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	common.ResponseSuccess(ctx, gin.H{
		"images": t,
	})
}
