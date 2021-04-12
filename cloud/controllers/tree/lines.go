package tree

import (
	"github.com/gin-gonic/gin"
	"github.com/mzz2017/grip/cloud/common"
	"github.com/mzz2017/grip/cloud/dao"
)

func GetLines(ctx *gin.Context) {
	type Result struct {
		ID          uint   `gorm:"column:id"`
		Name        string `gorm:"column:name"`
		CompanyID   uint   `gorm:"column:companyID"`
		CompanyName string `gorm:"column:companyName"`
		PowerID     uint   `gorm:"column:powerID"`
		PowerName   string `gorm:"column:powerName"`
	}
	rows, err := dao.DB().Raw("select l.id id,l.name name,c.id companyID,c.name companyName,p.id powerID,p.name powerName from companies c,powers p,`lines` l where l.power_id=p.id and l.company_id=c.id").Rows()
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	result := make([]*Result, 0)
	defer rows.Close()
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
