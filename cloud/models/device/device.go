package device

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/mzz2017/grip/cloud/dao"
)

type Device struct {
	gorm.Model
	Name         string `gorm:"not null;unique"`
	TowerID      int    `gorm:"not null"`
	Manufacture  string `gorm:"not null"`
	ModelName    string `gorm:"not null"`
	SerialNumber string `gorm:"not null"`
	Disabled     bool   `gorm:"not null;default:false"`
}

func (d Device) Insert() error {
	return dao.DB().Model(&Device{}).Create(&d).Error
}

func (d Device) Update() (err error) {
	var e Device
	if dao.DB().Model(&d).First(&e); e.CreatedAt.IsZero() {
		return errors.New("invalid ID")
	}
	d.Disabled = e.Disabled
	return dao.DB().Model(&d).Updates(d).Error
}

func (d Device) UpdateDisabled() (err error) {
	var cnt int
	if dao.DB().Model(&d).Count(&cnt); cnt == 0 {
		return errors.New("invalid ID")
	}
	return dao.DB().Model(&d).Update("disabled", d.Disabled).Error
}

func (d Device) Delete() error {
	return dao.DB().Unscoped().Delete(&d).Error
}

func Get(id uint) (o Device, err error) {
	db := dao.DB()
	err = db.Where("id=?", id).First(&o).Error
	return
}
