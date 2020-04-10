package line

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mzz2017/grip/cloud/common"
	"github.com/mzz2017/grip/cloud/dao"
	"github.com/mzz2017/grip/cloud/models/company"
	"github.com/mzz2017/grip/cloud/models/line"
	"github.com/mzz2017/grip/cloud/models/power"
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

type listResult struct {
	line.Line
	TowersNum int
}

func getLines(page, limit int) (result []*listResult, total int, err error) {
	rows, err := dao.DB().Raw("select l.*,count(t.id) towers_num " +
		"from `lines` l left join towers t on t.line_id=l.id").
		Group("l.id").Offset((page - 1) * limit).Limit(limit).Rows()
	if err != nil {
		return
	}
	result = make([]*listResult, 0)
	for rows.Next() {
		t := new(listResult)
		err := dao.DB().ScanRows(rows, &t)
		if err != nil {
			continue
		}
		result = append(result, t)
	}
	dao.DB().Table("lines").Count(&total)
	return
}
func getCompanies() (result []*company.Company, err error) {
	rows, err := dao.DB().Table("companies").Rows()
	if err != nil {
		return
	}
	result = make([]*company.Company, 0)
	for rows.Next() {
		t := new(company.Company)
		err := dao.DB().ScanRows(rows, &t)
		if err != nil {
			continue
		}
		result = append(result, t)
	}
	return
}
func getPowers() (result []*power.Power, err error) {
	rows, err := dao.DB().Table("powers").Rows()
	if err != nil {
		return
	}
	result = make([]*power.Power, 0)
	for rows.Next() {
		t := new(power.Power)
		err := dao.DB().ScanRows(rows, &t)
		if err != nil {
			continue
		}
		result = append(result, t)
	}
	return
}

func GetLines(ctx *gin.Context) {
	page, pageSize, err := getParams(ctx)
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	lines, total, err := getLines(page, pageSize)
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	companies, err := getCompanies()
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	powers, err := getPowers()
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	common.ResponseSuccess(ctx, gin.H{"lines": gin.H{"data": lines, "total": total}, "companies": companies, "powers": powers})
}
