package visualization

import (
	"github.com/gin-gonic/gin"
	"github.com/mzz2017/grip/cloud/common"
	"github.com/mzz2017/grip/cloud/dao"
	"github.com/mzz2017/grip/cloud/models/baseStation"
	"github.com/mzz2017/grip/cloud/models/tower"
)

func GetLocations(ctx *gin.Context) {
	rows, err := dao.DB().Table("base_stations").Rows()
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	result := make([]*baseStation.BaseStation, 0)
	for rows.Next() {
		t := new(baseStation.BaseStation)
		err := dao.DB().ScanRows(rows, &t)
		if err != nil {
			continue
		}
		result = append(result, t)
	}
	rows, err = dao.DB().Table("towers").Rows()
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	result2 := make([]*tower.Tower, 0)
	for rows.Next() {
		t := new(tower.Tower)
		err := dao.DB().ScanRows(rows, &t)
		if err != nil {
			continue
		}
		result2 = append(result2, t)
	}
	common.ResponseSuccess(ctx, gin.H{"baseStations": result, "towers": result2})
}
