package tower

import (
	"github.com/gin-gonic/gin"
	"github.com/mzz2017/grip/cloud/common"
	"github.com/mzz2017/grip/cloud/dao"
	"github.com/mzz2017/grip/cloud/models/image"
	"time"
)

func GetMarksThisMonth(ctx *gin.Context) {
	id := ctx.Param("id")
	date := ctx.DefaultQuery("date", time.Now().Format("2006-01"))
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		if t, err = time.Parse("2006-01", date); err != nil {
			common.ResponseError(ctx, err)
			return
		}
	}
	thisMonthBegin := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.UTC)
	nextMonthBegin := time.Date(t.Year(), (t.Month()+1)%12, 1, 0, 0, 0, 0, time.UTC)
	rows, err := dao.DB().Raw("select i.* from images i,devices d where i.created_at>=? and i.created_at<? and i.device_id=d.id and d.tower_id=? and i.mark!='safe'", thisMonthBegin, nextMonthBegin, id).Rows()
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}
	result := make([]*image.Image, 0)
	defer rows.Close()
	for rows.Next() {
		m := new(image.Image)
		_ = dao.DB().ScanRows(rows, &m)
		result = append(result, m)
	}
	common.ResponseSuccess(ctx, gin.H{"marks": result})
}
