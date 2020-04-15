package image

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/mzz2017/grip/cloud/dao"
)

type Image struct {
	gorm.Model
	DeviceID int     `gorm:"not null"`
	Mark     *string `gorm:"null" json:"omitempty"`
	Annotate *string `gorm:"null" json:"omitempty"`
	Filename string  `gorm:"not null;unique"`
}

func GetByID(id int) (img Image, err error) {
	err = dao.DB().Where("id=?", id).First(&img).Error
	return
}

func GetByFilename(fname string) (img Image, err error) {
	err = dao.DB().Where("filename=?", fname).First(&img).Error
	return
}

func GetByTowerIDAfter(towerID int, after int, limit int) (imgs []Image, err error) {
	db := dao.DB()
	rows, err := db.Where("tower_id=? and id>after", towerID, after).Limit(limit).Rows()
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var tmp Image
		db.ScanRows(rows, &tmp)
		imgs = append(imgs, tmp)
	}
	return
}

func (image Image) Insert() error {
	return dao.DB().Model(&Image{}).Create(&image).Error
}

func (image Image) FirstNotMarkAfter() (u Image, err error) {
	db := dao.DB().Begin()
	defer func() {
		if err != nil {
			db.Commit()
		} else {
			db.Rollback()
		}
	}()
	err = db.Model(&Image{}).Where("id>? and mark is null", image.ID).First(&u).Error
	if err != nil {
		return
	}
	err = u.Update()
	return
}
func (image Image) FirstAfter() (u Image, err error) {
	db := dao.DB().Begin()
	defer func() {
		if err != nil {
			db.Commit()
		} else {
			db.Rollback()
		}
	}()
	err = db.Model(&Image{}).Where("id>?", image.ID).First(&u).Error
	if err != nil {
		return
	}
	err = u.Update()
	return
}
func (image Image) FirstBefore() (u Image, err error) {
	err = dao.DB().Model(&Image{}).Where("id<? and marker=?", image.ID).Order("id DESC").First(&u).Error
	return
}

func (image Image) Update() (err error) {
	var cnt int
	if dao.DB().Model(&image).Count(&cnt); cnt == 0 {
		return errors.New("invalid ID")
	}
	return dao.DB().Model(&image).Updates(image).Error
}

func (image Image) ListAfter(num int) (images []*Image) {
	rows, err := dao.DB().Model(&Image{}).Where("id>?", image.ID).Limit(num).Rows()
	if err != nil {
		return
	}
	images = make([]*Image, 0, num)
	for rows.Next() {
		t := new(Image)
		err := dao.DB().ScanRows(rows, t)
		if err != nil {
			continue
		}
		images = append(images, t)
	}
	return
}
