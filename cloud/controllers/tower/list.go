package tower

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mzz2017/grip/cloud/common"
	"github.com/mzz2017/grip/cloud/dao"
	"github.com/mzz2017/grip/cloud/models/baseStation"
	"github.com/mzz2017/grip/cloud/models/tower"
	"strconv"
)

func getParams(ctx *gin.Context) (page, pageSize int, lineID int, err error) {
	_page := ctx.DefaultQuery("page", "1")
	_pageSize := ctx.DefaultQuery("pageSize", "10")
	page, err = strconv.Atoi(_page)
	if _page == "" || err != nil {
		err = errors.New("invalid page")
		return
	}
	pageSize, err = strconv.Atoi(_pageSize)
	if _pageSize == "" || err != nil {
		err = errors.New("invalid pageSize")
		return
	}
	if page < 1 || pageSize < 0 {
		err = errors.New("invalid params")
		return
	}
	_lineID := ctx.Param("lineID")
	lineID, err = strconv.Atoi(_lineID)
	if _lineID == "" || err != nil {
		err = errors.New("invalid lineID")
	}
	return
}

func getTowers(page, limit int, lineID uint) (result []*tower.Tower, total int, err error) {
	rows, err := dao.DB().Table("towers").Where("line_id=?", lineID).Offset((page - 1) * limit).Limit(limit).Rows()
	if err != nil {
		return
	}
	result = make([]*tower.Tower, 0)
	for rows.Next() {
		t := new(tower.Tower)
		err := dao.DB().ScanRows(rows, &t)
		if err != nil {
			continue
		}
		result = append(result, t)
	}
	dao.DB().Table("towers").Count(&total)
	return
}
func getBaseStations() (result []*baseStation.BaseStation, err error) {
	rows, err := dao.DB().Table("base_stations").Rows()
	if err != nil {
		return
	}
	result = make([]*baseStation.BaseStation, 0)
	for rows.Next() {
		t := new(baseStation.BaseStation)
		err := dao.DB().ScanRows(rows, &t)
		if err != nil {
			continue
		}
		result = append(result, t)
	}
	return
}

func GetTowers(ctx *gin.Context) {
	page, pageSize, lineID, err := getParams(ctx)
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	users, total, err := getTowers(page, pageSize, uint(lineID))
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	baseStations, err := getBaseStations()
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	common.ResponseSuccess(ctx, gin.H{"towers": gin.H{"data": users, "total": total}, "baseStations": baseStations})
}
