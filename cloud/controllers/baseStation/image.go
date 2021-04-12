package baseStation

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mzz2017/grip/cloud/models/baseStation"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
)

func GetImage(ctx *gin.Context) {
	_id := ctx.Param("id")
	id, err := strconv.Atoi(_id)
	if err != nil {
		ctx.Status(400)
		ctx.Writer.WriteString(err.Error())
		return
	}
	o, err := baseStation.Get(uint(id))
	if err != nil {
		ctx.Status(400)
		ctx.Writer.WriteString("BaseStation error: " + err.Error())
		return
	}
	filename := ctx.Param("filename")
	u, _ := url.Parse(fmt.Sprintf("http://%v:3000", o.IP))
	proxy := httputil.NewSingleHostReverseProxy(u)
	//url重写，以正确代理转发
	u, _ = url.Parse(fmt.Sprintf("http://%v/api/image/%v", ctx.Request.Host, filename))
	ctx.Request.URL = u
	proxy.ErrorHandler = func(writer http.ResponseWriter, request *http.Request, err error) {
		ctx.Status(400)
		ctx.Writer.WriteString(err.Error())
	}
	proxy.ServeHTTP(ctx.Writer, ctx.Request)
}
