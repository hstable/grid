package line

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/mzz2017/grip/cloud/common"
	"github.com/mzz2017/grip/cloud/models/line"
	"strconv"
)

func PostLine(ctx *gin.Context) {
	var l line.Line
	err := ctx.BindJSON(&l)
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	l.Model = gorm.Model{}
	err = l.Insert()
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	common.ResponseSuccess(ctx, nil)
}

func Put(ctx *gin.Context) {
	var t line.Line
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
	o, err := line.Get(uint(id))
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	common.ResponseSuccess(ctx, gin.H{
		"line": o,
	})
}
