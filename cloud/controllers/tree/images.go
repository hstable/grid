package tree

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mzz2017/grip/cloud/common"
	"github.com/mzz2017/grip/cloud/models/image"
	"strconv"
	"time"
)

func getParams(ctx *gin.Context) (towerID, after, limit int, beginTime, endTime time.Time, err error) {
	_towerID := ctx.Param("towerID")
	_after := ctx.DefaultQuery("after", "0")
	_limit := ctx.DefaultQuery("limit", "4")
	_beginTime := ctx.Query("beginTime")
	_endTime := ctx.Query("endTime")
	towerID, err = strconv.Atoi(_towerID)
	if _towerID == "" || err != nil {
		err = errors.New("invalid towerID")
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
	if _beginTime != "" {
		beginTime, err = time.Parse(time.RFC3339, _beginTime)
		if err != nil {
			err = errors.New("invalid beginTime")
			return
		}
	} else {
		beginTime = time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	}
	if _endTime != "" {
		endTime, err = time.Parse(time.RFC3339, _endTime)
		if err != nil {
			err = errors.New("invalid endTime")
			return
		}
	} else {
		endTime = time.Now()
	}
	return
}

func GetImages(ctx *gin.Context) {
	towerID, after, limit, beginTime, endTime, err := getParams(ctx)
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	t, err := image.GetByTowerIDAfter(towerID, after, limit, beginTime, endTime)
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	common.ResponseSuccess(ctx, gin.H{
		"images": t,
	})
}
