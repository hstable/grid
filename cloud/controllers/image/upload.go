package image

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mzz2017/grip/cloud/common"
	"github.com/mzz2017/grip/cloud/models/image"
	"io/ioutil"
	"strconv"
	"strings"
)

func getParams(ctx *gin.Context) (mark string, deviceID int, image []byte, err error) {
	mark = ctx.Request.FormValue("mark")
	mark = strings.TrimSpace(mark)
	if mark == "" {
		return "", 0, nil, errors.New("invalid params")
	}
	//FIXME: 没有验证该device是否属于该基站
	_deviceID := ctx.Request.FormValue("deviceID")
	deviceID, err = strconv.Atoi(_deviceID)
	if err != nil {
		return "", 0, nil, errors.New("invalid params")
	}
	file, header, err := ctx.Request.FormFile("image")
	if err != nil {
		return "", 0, nil, errors.New("invalid params: " + err.Error())
	}
	if header == nil {
		return "", 0, nil, errors.New("invalid params")
	}
	image, _ = ioutil.ReadAll(file)
	return
}

func PostUpload(ctx *gin.Context) {
	mark, deviceID, img, err := getParams(ctx)
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	err = image.Image{
		DeviceID: deviceID,
		Mark:     &mark,
	}.Insert(img)
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	common.ResponseSuccess(ctx, nil)
}


