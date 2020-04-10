package power

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/mzz2017/grip/cloud/dao"
)

type Power struct {
	gorm.Model
	Name string `gorm:"not null;unique"`
}

func (p Power) Insert() error {
	return dao.DB().Model(&Power{}).Create(&p).Error
}

func (p Power) Update() (err error) {
	var cnt int
	if dao.DB().Model(&p).Count(&cnt); cnt == 0 {
		return errors.New("invalid ID")
	}
	return dao.DB().Model(&p).Updates(p).Error
}
func Get(id uint) (o Power, err error) {
	db := dao.DB()
	err = db.Where("id=?", id).First(&o).Error
	return
}
func (p Power) Delete() error {
	return dao.DB().Unscoped().Delete(&p).Error
}