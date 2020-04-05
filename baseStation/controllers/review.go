package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mzz2017/grip/basestation/api"
	"github.com/mzz2017/grip/basestation/common"
	"github.com/mzz2017/grip/basestation/service/image"
	"io/ioutil"
	"strconv"
)

func PostReview(ctx *gin.Context) {
	file, header, err := ctx.Request.FormFile("image")
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	if header == nil {
		common.ResponseError(ctx, errors.New("invalid params"))
		return
	}
	_deviceID := ctx.Request.FormValue("deviceID")
	var deviceID int
	//FIXME: 没有验证设备与IP对应关系
	if deviceID, err = strconv.Atoi(_deviceID); err != nil {
		common.ResponseError(ctx, errors.New("unrecognized device"))
		return
	}
	b, _ := ioutil.ReadAll(file)
	if !common.IsImage(b) {
		common.ResponseError(ctx, errors.New("unrecognized image format"))
		return
	}
	level := image.ValidateImage(b)
	if level != image.Safe {
		err = api.UploadImage(b, level, deviceID)
		if err != nil {
			return
		}

	}
	common.ResponseSuccess(ctx, gin.H{
		"level": level,
	})
}
