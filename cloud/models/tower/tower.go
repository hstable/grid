package tower

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/mzz2017/grip/cloud/dao"
	"github.com/mzz2017/grip/cloud/dataStructure/location"
)

type Tower struct {
	gorm.Model
	LineID        int               `gorm:"not null"`
	BaseStationID int               `gorm:"not null"`
	Name          string            `gorm:"not null"`
	Location      location.Location `gorm:"not null"` //lat,lon
}

func GetByLineID(lineID int) (towers []Tower, err error) {
	db := dao.DB()
	rows, err := db.Where("line_id=?", lineID).Rows()
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var tmp Tower
		db.ScanRows(rows, &tmp)
		towers = append(towers, tmp)
	}
	return
}
func Get(id uint) (tower Tower, err error) {
	db := dao.DB()
	err = db.Where("id=?", id).First(&tower).Error
	return
}
func (t Tower) Update() (err error) {
	var cnt int
	if dao.DB().Model(&t).Count(&cnt); cnt == 0 {
		return errors.New("invalid ID")
	}
	return dao.DB().Model(&t).Updates(t).Error
}
func (t Tower) Insert() error {
	return dao.DB().Model(&Tower{}).Create(&t).Error
}
