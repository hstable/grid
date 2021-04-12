package visualization

import (
	"github.com/gin-gonic/gin"
	"github.com/mzz2017/grip/cloud/common"
	"github.com/mzz2017/grip/cloud/dao"
)

func GetMarks(ctx *gin.Context) {
	type Result struct {
		Mark string
		Count     int
	}
	rows, err := dao.DB().Raw("select mark,count(id) count from images group by mark").Rows()
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
	common.ResponseSuccess(ctx, gin.H{"result": result})
}
