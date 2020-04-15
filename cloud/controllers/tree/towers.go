package tree

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mzz2017/grip/cloud/common"
	"github.com/mzz2017/grip/cloud/models/tower"
	"strconv"
)

func GetTowers(ctx *gin.Context) {
	_lineID := ctx.Param("lineID")
	lineID, err := strconv.Atoi(_lineID)
	if _lineID == "" || err != nil {
		common.ResponseError(ctx, errors.New("invalid lineID"))
		return
	}
	t, err := tower.GetByLineID(lineID)
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	common.ResponseSuccess(ctx, gin.H{
		"towers": t,
	})
}
