package power

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/mzz2017/grip/cloud/common"
	"github.com/mzz2017/grip/cloud/models/power"
	"strconv"
)

func PostPower(ctx *gin.Context) {
	var p power.Power
	err := ctx.BindJSON(&p)
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	p.Model = gorm.Model{}
	err = p.Insert()
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	common.ResponseSuccess(ctx, nil)
}

func Put(ctx *gin.Context) {
	var t power.Power
	err := ctx.BindJSON(&t)
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	err = t.Update()
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	common.ResponseSuccess(ctx, nil)
}

func Get(ctx *gin.Context) {
	_id, ok := ctx.GetQuery("id")
	if !ok {
		common.ResponseError(ctx, errors.New("invalid params"))
		return
	}
	id, err := strconv.Atoi(_id)
	if err != nil {
		common.ResponseError(ctx, errors.New("invalid params"))
		return
	}
	o, err := power.Get(uint(id))
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	common.ResponseSuccess(ctx, gin.H{
		"power": o,
	})
}
