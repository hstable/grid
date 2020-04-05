package device

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/mzz2017/grip/cloud/common"
	"github.com/mzz2017/grip/cloud/models/device"
	"strconv"
)

func PostDevice(ctx *gin.Context) {
	var d device.Device
	err := ctx.BindJSON(&d)
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	d.Model = gorm.Model{}
	err = d.Insert()
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	common.ResponseSuccess(ctx, nil)
}

func PuDevice(ctx *gin.Context) {
	var t device.Device
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
	o, err := device.Get(uint(id))
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	common.ResponseSuccess(ctx, gin.H{
		"device": o,
	})
}
