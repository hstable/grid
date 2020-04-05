package device

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/mzz2017/grip/cloud/dao"
)

type Device struct {
	gorm.Model
	TowerID     int    `gorm:"not null"`
	Manufacture string `gorm:"not null"`
	ModelName   string `gorm:"not null"`
	Disable     bool   `gorm:"not null;default:false"`
}

func (d Device) Insert() error {
	return dao.DB().Model(&Device{}).Create(&d).Error
}

func (d Device) Update() (err error) {
	var cnt int
	if dao.DB().Model(&d).Count(&cnt); cnt == 0 {
		return errors.New("invalid ID")
	}
	return dao.DB().Model(&d).Updates(d).Error
}

func Get(id uint) (o Device, err error) {
	db := dao.DB()
	err = db.Where("id=?", id).First(&o).Error
	return
}