package baseStation

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/mzz2017/grip/cloud/common"
	"github.com/mzz2017/grip/cloud/models/baseStation"
	"strconv"
)

func Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	err = baseStation.BaseStation{Model: gorm.Model{ID: uint(id)}}.Delete()
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	common.Response(ctx, common.SUCCESS, nil)
}
func Post(ctx *gin.Context) {
	var p baseStation.BaseStation
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
	var t baseStation.BaseStation
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

func PutOnline(ctx *gin.Context) {
	var t baseStation.BaseStation
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
	err = t.UpdateOnline()
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
	o, err := baseStation.Get(uint(id))
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	common.ResponseSuccess(ctx, gin.H{
		"baseStation": o,
	})
}
