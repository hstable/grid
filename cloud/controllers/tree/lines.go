package tree

import (
	"github.com/gin-gonic/gin"
	"github.com/mzz2017/grip/cloud/common"
	"github.com/mzz2017/grip/cloud/dao"
)

func Get(ctx *gin.Context) {
	type Result struct {
		ID          uint   `json:"id"`
		Name        string `json:"name"`
		CompanyID   uint   `json:"companyID"`
		CompanyName string `json:"companyName"`
		PowerID     uint   `json:"powerID"`
		PowerName   string `json:"powerName"`
	}
	rows, err := dao.DB().Raw("select (l.id id,l.name name,c.id companyID,c.name companyName,p.id powerID,p.name powerName) from companies c,powers p,line l where l.power_id=p.id and l.company_id=p.id").Rows()
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	result := make([]*Result, 0)
	for rows.Next() {
		t := new(Result)
		err := dao.DB().ScanRows(rows, &t)
		if err != nil {
			continue
		}
		result = append(result, t)
	}
	common.ResponseSuccess(ctx, gin.H{"lines": result})
}
