package tower

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/mzz2017/grip/cloud/common"
	"github.com/mzz2017/grip/cloud/models/tower"
	"strconv"
)

func PostTower(ctx *gin.Context) {
	var t tower.Tower
	err := ctx.BindJSON(&t)
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	t.Model = gorm.Model{}
	err = t.Insert()
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	common.ResponseSuccess(ctx, nil)
}
func Put(ctx *gin.Context) {
	var t tower.Tower
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
	o, err := tower.Get(uint(id))
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	common.ResponseSuccess(ctx, gin.H{
		"tower": o,
	})
}
