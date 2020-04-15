package visualization

import (
	"github.com/gin-gonic/gin"
	"github.com/mzz2017/grip/cloud/common"
	"github.com/mzz2017/grip/cloud/dao"
)

func GetPowerCounts(ctx *gin.Context) {
	type Result struct {
		PowerName string
		Count     int
	}
	rows, err := dao.DB().Raw("select p.name power_name,count(p.id) count from towers t,`lines` l,powers p where t.line_id=l.id and l.power_id=p.id group by power_id").Rows()
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
