package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mzz2017/grip/basestation/common"
	"github.com/mzz2017/grip/basestation/config"
	"github.com/mzz2017/grip/basestation/global"
	"github.com/mzz2017/grip/basestation/models"
	"github.com/mzz2017/hcnetsdk-go"
	"io/ioutil"
	"os"
	"path"
)

func PostPicture() {

}

func PostDevices(ctx *gin.Context) {
	var params struct {
		Devices models.DeviceNetworks `json:"devices"`
	}
	err := ctx.BindJSON(&params)
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	csv := []byte(params.Devices.ToCSV())
	err = ioutil.WriteFile(
		path.Join(config.Get().SrslteConfigDir, "user_db.csv"),
		csv,
		os.FileMode(0644),
	)
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	global.LoadMapIPDevice(csv)
	common.ResponseSuccess(ctx, nil)
}

func GetNewPhoto(ctx *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			ctx.Status(400)
			ctx.Writer.Write([]byte("!ERROR"))
			ctx.Writer.Write([]byte{0})
			ctx.Writer.Write([]byte(err.Error()))
		}
	}()
	ip := ctx.Param("ip")
	//TODO: cache
	sdk := new(hcnetsdk.HCNetSDK)
	sdk.Init()
	err = sdk.Login(ip, 8000, "admin", "HikQPUWBS")
	if err != nil {
		return
	}
	defer func() {
		ok := sdk.Logout()
		if !ok && err == nil {
			err = errors.New("fail in logout")
		}
	}()
	err, b := sdk.CaptureJPEGPictureNew(&hcnetsdk.JPEGParam{
		PicSize:    9,
		PicQuality: 0,
	})
	if err != nil {
		return
	}
	ctx.Status(200)
	//TODO: mark
	mark := "safe"
	ctx.Writer.Write([]byte(mark))
	ctx.Writer.Write([]byte{0})
	ctx.Writer.Write(b)
}
