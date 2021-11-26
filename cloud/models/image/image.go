package image

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/mzz2017/grip/cloud/common"
	"github.com/mzz2017/grip/cloud/common/crypto"
	"github.com/mzz2017/grip/cloud/config"
	"github.com/mzz2017/grip/cloud/dao"
	"io/ioutil"
	"path"
	"path/filepath"
	"time"
)

type Image struct {
	gorm.Model
	DeviceID int     `gorm:"not null"`
	Mark     *string `gorm:"null" json:"Mark,omitempty"`
	Annotate *string `gorm:"null" json:"Annotate,omitempty"`
	Filename string  `gorm:"not null;unique"`
	Type     string  `gorm:"default:'';not null"`
}

func GetByID(id int) (img Image, err error) {
	err = dao.DB().Where("id=?", id).First(&img).Error
	return
}

func GetByFilename(fname string) (img Image, err error) {
	err = dao.DB().Where("filename=?", fname).First(&img).Error
	return
}

func GetByTowerIDAfter(towerID int, after int, limit int, beginTime, endTime time.Time, typ string) (imgs []Image, err error) {
	if endTime.IsZero() {
		endTime = time.Now()
	}
	db := dao.DB()
	rows, err := db.Raw("select i.* from images i,devices d where device_id=d.id and tower_id=? and i.id<? and i.created_at>=? and i.created_at<? and type=?", towerID, after, beginTime, endTime, typ).Limit(limit).Order("id desc").Rows()
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

func writeToFile(image []byte) (fpath string, err error) {
	suffix := common.GetImageSuffix(image)
	if suffix == "" {
		return "", errors.New("unrecognized image format")
	}
	fname := crypto.HashWithSalt(image, config.Get().Secret) + suffix
	fpath = path.Join(config.Get().AssetDir, fname)
	return fpath, ioutil.WriteFile(fpath, image, 0700)
}

func (image Image) Insert(b []byte) error {
	fpath, err := writeToFile(b)
	if err != nil {
		return err
	}
	image.Filename = filepath.Base(fpath)
	return dao.DB().Model(&Image{}).Create(&image).Error
}

func (image Image) FirstNotMarkAfter() (u Image, err error) {
	db := dao.DB().Begin()
	defer func() {
		if err != nil {
			db.Rollback()
		} else {
			db.Commit()
		}
	}()
	err = db.Model(&Image{}).Where("id>? and mark is null and typ=?", image.ID, image.Type).First(&u).Error
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
			db.Rollback()
		} else {
			db.Commit()
		}
	}()
	err = db.Model(&Image{}).Where("id>? and typ=?", image.ID, image.Type).First(&u).Error
	if err != nil {
		return
	}
	err = u.Update()
	return
}
func (image Image) FirstBefore() (u Image, err error) {
	err = dao.DB().Model(&Image{}).Where("id<? and marker=? and typ=?", image.ID, image.Type).Order("id DESC").First(&u).Error
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
	rows, err := dao.DB().Model(&Image{}).Where("id>? and typ=?", image.ID, image.Type).Limit(num).Rows()
	if err != nil {
		return
	}
	images = make([]*Image, 0, num)
	defer rows.Close()
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
