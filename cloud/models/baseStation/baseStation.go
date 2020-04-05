package baseStation

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/mzz2017/grip/cloud/dao"
	"github.com/mzz2017/grip/cloud/dataStructure/location"
)

type BaseStation struct {
	gorm.Model
	IP       int               `gorm:"not null"`
	Location location.Location `gorm:"not null"`
	Online   bool              `gorm:"not null"`
}

func (b BaseStation) Insert() error {
	return dao.DB().Model(&BaseStation{}).Create(&b).Error
}

func (b BaseStation) Update() (err error) {
	var cnt int
	if dao.DB().Model(&b).Count(&cnt); cnt == 0 {
		return errors.New("invalid ID")
	}
	return dao.DB().Model(&b).Updates(b).Error
}

func Get(id uint) (o BaseStation, err error) {
	db := dao.DB()
	err = db.Where("id=?", id).First(&o).Error
	return
}
