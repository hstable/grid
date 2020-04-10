package device

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/mzz2017/grip/cloud/common"
	"github.com/mzz2017/grip/cloud/models/device"
	"strconv"
)

func Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	err = device.Device{Model: gorm.Model{ID: uint(id)}}.Delete()
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	common.Response(ctx, common.SUCCESS, nil)
}
func Post(ctx *gin.Context) {
	var d device.Device
	err := ctx.BindJSON(&d)
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	_towerID := ctx.Param("towerID")
	towerID, err := strconv.Atoi(_towerID)
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	d.TowerID = towerID
	d.Model = gorm.Model{}
	err = d.Insert()
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	common.ResponseSuccess(ctx, nil)
}

func Put(ctx *gin.Context) {
	var t device.Device
	err := ctx.BindJSON(&t)
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	t.ID = uint(id)
	err = t.Update()
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	common.ResponseSuccess(ctx, nil)
}

func PutDisabled(ctx *gin.Context) {
	var t device.Device
	err := ctx.BindJSON(&t)
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	t.ID = uint(id)
	err = t.UpdateDisabled()
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
