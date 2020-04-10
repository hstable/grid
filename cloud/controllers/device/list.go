package device

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mzz2017/grip/cloud/common"
	"github.com/mzz2017/grip/cloud/dao"
	"github.com/mzz2017/grip/cloud/models/device"
	"strconv"
)

func getParams(ctx *gin.Context) (page, pageSize int, towerID int, err error) {
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
	_towerID := ctx.Param("towerID")
	towerID, err = strconv.Atoi(_towerID)
	if _towerID == "" || err != nil {
		err = errors.New("invalid towerID")
	}
	return
}

func getDevices(page, limit int, towerID uint) (result []*device.Device, total int, err error) {
	rows, err := dao.DB().Table("devices").Where("tower_id=?", towerID).Offset((page - 1) * limit).Limit(limit).Rows()
	if err != nil {
		return
	}
	result = make([]*device.Device, 0)
	for rows.Next() {
		t := new(device.Device)
		err := dao.DB().ScanRows(rows, &t)
		if err != nil {
			continue
		}
		result = append(result, t)
	}
	dao.DB().Table("devices").Where("tower_id=?", towerID).Count(&total)
	return
}

func GetDevices(ctx *gin.Context) {
	page, pageSize, towerID, err := getParams(ctx)
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	users, total, err := getDevices(page, pageSize, uint(towerID))
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	common.ResponseSuccess(ctx, gin.H{"devices": gin.H{"data": users, "total": total}})
}
