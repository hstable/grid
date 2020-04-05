package company

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/mzz2017/grip/cloud/dao"
)

type Company struct {
	gorm.Model
	Name string `gorm:"not null"`
}

func (c Company) Insert() error {
	return dao.DB().Model(&Company{}).Create(&c).Error
}

func (c Company) Update() (err error) {
	var cnt int
	if dao.DB().Model(&c).Count(&cnt); cnt == 0 {
		return errors.New("invalid ID")
	}
	return dao.DB().Model(&c).Updates(c).Error
}

func Get(id uint) (o Company, err error) {
	db := dao.DB()
	err = db.Where("id=?", id).First(&o).Error
	return
}