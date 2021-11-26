package visualization

import (
	"github.com/gin-gonic/gin"
	"github.com/mzz2017/grip/cloud/common"
	"github.com/mzz2017/grip/cloud/dao"
	"strings"
)

func GetTypeCounts(ctx *gin.Context) {
	const delimiter = " + "
	type SqlResult struct {
		Annotate string
	}
	type Result struct {
		Value int    `json:"value"`
		Type  string `json:"type"`
		Name  string `json:"name"`
	}
	rows, err := dao.DB().Raw("select annotate from images where mark != 'safe'").Rows()
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	result := make([]Result, 0)
	defer rows.Close()
	cnt := make(map[string]int)
	for rows.Next() {
		t := new(SqlResult)
		err := dao.DB().ScanRows(rows, &t)
		if err != nil {
			continue
		}
		fields := strings.Split(t.Annotate, delimiter)
		for _, field := range fields {
			cnt[field]++
		}
	}
	for field := range cnt {
		result = append(result, Result{
			Value: cnt[field],
			Type:  "施工", //FIXME
			Name:  field,
		})
	}
	common.ResponseSuccess(ctx, gin.H{"result": result})
}
