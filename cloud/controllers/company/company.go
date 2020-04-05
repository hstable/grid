package company

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/mzz2017/grip/cloud/common"
	"github.com/mzz2017/grip/cloud/models/company"
	"strconv"
)

func Post(ctx *gin.Context) {
	var c company.Company
	err := ctx.BindJSON(&c)
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	c.Model = gorm.Model{}
	err = c.Insert()
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	common.ResponseSuccess(ctx, nil)
}

func Put(ctx *gin.Context) {
	var t company.Company
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
	o, err := company.Get(uint(id))
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	common.ResponseSuccess(ctx, gin.H{
		"company": o,
	})
}
