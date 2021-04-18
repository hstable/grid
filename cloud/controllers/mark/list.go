package mark

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mzz2017/grip/cloud/common"
	"github.com/mzz2017/grip/cloud/dao"
	"github.com/mzz2017/grip/cloud/models/image"
	"strconv"
)

func getParams(ctx *gin.Context) (page, pageSize int, err error) {
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
	}
	return
}

func GetMarksRisky(ctx *gin.Context) {
	type Result struct {
		image.Image
		TowerName       string
		TowerID         uint
		LineName        string
		LineID          uint
		PowerName       string
		PowerID         uint
		BaseStationID   uint
		BaseStationName string
	}
	var err error
	defer func() {
		if err != nil {
			common.ResponseError(ctx, err)
		}
	}()
	page, pageSize, err := getParams(ctx)
	if err != nil {
		return
	}
	rows, err := dao.DB().Raw(`
select i.*, t.id tower_id, t.name tower_name, l.id line_id, l.name line_name, p.id power_id, p.name power_name,t.base_station_id,b.name base_station_name
from images i
         inner join devices d
         inner join towers t
         inner join ` + "`lines`" + ` l
         inner join powers p
         inner join base_stations b
on i.device_id = d.id and d.tower_id = t.id and t.line_id = l.id and l.power_id = p.id and t.base_station_id=b.id
where i.mark != 'safe'
`).Offset((page - 1) * pageSize).Limit(pageSize).Order("i.created_at desc").Rows()
	if err != nil {
		return
	}
	result := make([]*Result, 0)
	defer rows.Close()
	for rows.Next() {
		m := new(Result)
		_ = dao.DB().ScanRows(rows, &m)
		result = append(result, m)
	}
	var total int
	dao.DB().Table("images").Where("mark!='safe'").Count(&total)
	common.ResponseSuccess(ctx, gin.H{"marks": gin.H{"data": result, "total": total}})
}
