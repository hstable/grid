package edgeServer

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mzz2017/grip/cloud/common"
	"github.com/mzz2017/grip/cloud/models/baseStation"
	"net/http/httputil"
	"net/url"
	"strconv"
)

func GetImage(ctx *gin.Context) {
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
	filename := ctx.Param("filename")
	u, _ := url.Parse(fmt.Sprintf("http://%v:3000", o.IP))
	proxy := httputil.NewSingleHostReverseProxy(u)
	//url重写，以正确代理转发
	u, _ = url.Parse(fmt.Sprintf("http://%v/api/image/%v", ctx.Request.Host, filename))
	ctx.Request.URL = u
	proxy.ServeHTTP(ctx.Writer, ctx.Request)
}
