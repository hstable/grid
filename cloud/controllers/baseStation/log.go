package baseStation

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mzz2017/grip/cloud/common"
	"github.com/mzz2017/grip/cloud/models/baseStation"
	"io/ioutil"
	"net/http"
	"strconv"
)

func GetLog(ctx *gin.Context) {
	_id := ctx.Param("id")
	id, err := strconv.Atoi(_id)
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	o, err := baseStation.Get(uint(id))
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	resp, err := http.Get(fmt.Sprintf("http://%v:3000/api/log", o.IP))
	if err != nil {
		err = fmt.Errorf("与基站通信时出现错误: %v", err.Error())
		if o.Online {
			o.Online = false
			_ = o.UpdateOnline()
		}
		common.ResponseError(ctx, err)
		return
	} else {
		if !o.Online {
			o.Online = true
			_ = o.UpdateOnline()
		}
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		common.ResponseError(ctx, errors.New("与边缘服务器通信失败:"+err.Error()))
		return
	}
	defer resp.Body.Close()
	ctx.Status(200)
	ctx.Writer.Write(b)
}
