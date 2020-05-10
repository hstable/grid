package device

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/mzz2017/grip/cloud/common"
	"github.com/mzz2017/grip/cloud/models/device"
	image2 "github.com/mzz2017/grip/cloud/models/image"
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
	var err error
	defer func() {
		if err != nil {
			common.ResponseError(ctx, err)
		}
	}()
	var d device.Device
	err = ctx.BindJSON(&d)
	if err != nil {
		return
	}
	_towerID := ctx.Param("towerID")
	towerID, err := strconv.Atoi(_towerID)
	if err != nil {
		return
	}
	d.TowerID = towerID
	d.Model = gorm.Model{}
	err = d.Insert()
	if err != nil {
		return
	}
	bs, err := d.BaseStation()
	if err != nil {
		return
	}
	err = bs.PushDevices()
	if err != nil {
		return
	}
	common.ResponseSuccess(ctx, nil)
}

func Put(ctx *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			common.ResponseError(ctx, err)
		}
	}()
	var d device.Device
	err = ctx.BindJSON(&d)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return
	}
	d.ID = uint(id)
	err = d.Update()
	if err != nil {
		return
	}
	bs, err := d.BaseStation()
	if err != nil {
		return
	}
	err = bs.PushDevices()
	if err != nil {
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
func GetNewPhoto(ctx *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			err = fmt.Errorf("GetNewPhoto: %v", err.Error())
			ctx.Status(400)
			ctx.Writer.Write([]byte(err.Error()))
		}
	}()
	_deviceID := ctx.Param("id")
	deviceID, err := strconv.Atoi(_deviceID)
	if err != nil {
		return
	}
	img, mark, err := device.Device{
		Model: gorm.Model{
			ID: uint(deviceID),
		},
	}.Get().NewPhoto()
	if err != nil {
		return
	}
	err = image2.Image{
		DeviceID: deviceID,
		Mark:     &mark,
	}.Insert(img)
	ctx.Status(200)
	ctx.Writer.Write(img)
}
