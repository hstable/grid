package user

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mzz2017/grip/cloud/common"
	"github.com/mzz2017/grip/cloud/dao"
	"github.com/mzz2017/grip/cloud/models/company"
	"github.com/mzz2017/grip/cloud/models/user"
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

func getUsers(page, limit int) (result []*user.User, total int, err error) {
	rows, err := dao.DB().Table("users").Offset((page - 1) * limit).Limit(limit).Rows()
	if err != nil {
		return
	}
	result = make([]*user.User, 0)
	for rows.Next() {
		t := new(user.User)
		err := dao.DB().ScanRows(rows, &t)
		if err != nil {
			continue
		}
		t.Password = "$__hidden_PASSWORD"
		result = append(result, t)
	}
	dao.DB().Table("users").Count(&total)
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

func GetUsers(ctx *gin.Context) {
	page, pageSize, err := getParams(ctx)
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	users, total, err := getUsers(page, pageSize)
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	companies, err := getCompanies()
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	common.ResponseSuccess(ctx, gin.H{"users": gin.H{"data": users, "total": total}, "companies": companies})
}
