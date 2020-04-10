package line

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/mzz2017/grip/cloud/dao"
)

type Line struct {
	gorm.Model
	Name      string `gorm:"not null;unique"`
	PowerID   int
	CompanyID int
}

func (l Line) Insert() error {
	return dao.DB().Model(&Line{}).Create(&l).Error
}

func (l Line) Update() (err error) {
	var cnt int
	if dao.DB().Model(&l).Count(&cnt); cnt == 0 {
		return errors.New("invalid ID")
	}
	return dao.DB().Model(&l).Updates(l).Error
}
func Get(id uint) (o Line, err error) {
	db := dao.DB()
	err = db.Where("id=?", id).First(&o).Error
	return
}
func (l Line) Delete() error {
	return dao.DB().Unscoped().Delete(&l).Error
}