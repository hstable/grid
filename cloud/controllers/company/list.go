package company

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mzz2017/grip/cloud/common"
	"github.com/mzz2017/grip/cloud/dao"
	"github.com/mzz2017/grip/cloud/models/company"
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
func getCompanies(page, limit int) (result []*company.Company, total int, err error) {
	rows, err := dao.DB().Table("companies").Offset((page - 1) * limit).Limit(limit).Rows()
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
	dao.DB().Table("companies").Count(&total)
	return
}
func GetCompanies(ctx *gin.Context) {
	page, pageSize, err := getParams(ctx)
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	companies, total, err := getCompanies(page, pageSize)
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	common.ResponseSuccess(ctx, gin.H{"companies": gin.H{"data": companies, "total": total}})
}
