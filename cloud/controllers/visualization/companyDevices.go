package visualization

import (
	"github.com/gin-gonic/gin"
	"github.com/mzz2017/grip/cloud/common"
	"github.com/mzz2017/grip/cloud/dao"
)

func GetCompanyDevices(ctx *gin.Context) {
	type Result struct {
		CompanyName string
		Count     int
	}
	rows, err := dao.DB().Raw("select c.name company_name,count(d.id) count from companies c inner join devices d inner join towers t inner join `lines` l on c.id = l.company_id and t.line_id=l.id and t.id=d.tower_id group by c.name").Rows()
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
	common.ResponseSuccess(ctx, gin.H{"result": result})
}
