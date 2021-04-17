package device

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/mzz2017/grip/cloud/dao"
	"github.com/mzz2017/grip/cloud/models/baseStation"
	"github.com/mzz2017/grip/cloud/models/deviceNetwork"
	"io/ioutil"
	"net/http"
)

type Device struct {
	gorm.Model
	Name         string `gorm:"not null"`
	TowerID      int    `gorm:"not null"`
	Manufacture  string `gorm:"not null"`
	ModelName    string `gorm:"not null"`
	SerialNumber string `gorm:"not null"`
	Disabled     bool   `gorm:"not null;default:false"`
	Network      deviceNetwork.DeviceNetwork
	NetworkID    uint `gorm:"not null"`
}

func (d Device) BaseStation() (*baseStation.BaseStation, error) {
	db := dao.DB()
	var b baseStation.BaseStation
	err := db.Unscoped().Raw("select base_stations.* from base_stations,towers,devices where devices.id=? and devices.tower_id=towers.id and towers.base_station_id=base_stations.id", d.ID).First(&b).Error
	return &b, err
}

func (d *Device) Insert() (err error) {
	db := dao.DB()
	af := db.Begin()
	defer func() {
		if err != nil {
			af.Rollback()
		} else {
			af.Commit()
		}
	}()
	d.Network.Name = d.Name
	err = af.Model(&deviceNetwork.DeviceNetwork{}).Create(&d.Network).Error
	if err != nil {
		return
	}
	d.NetworkID = d.Network.ID
	return af.Model(&Device{}).Create(&d).Error
}

func (d *Device) Update() (err error) {
	var e Device
	if dao.DB().Model(&d).First(&e); e.CreatedAt.IsZero() {
		return errors.New("invalid ID")
	}
	d.Disabled = e.Disabled
	db := dao.DB()
	af := db.Begin()
	defer func() {
		if err != nil {
			af.Rollback()
		} else {
			af.Commit()
		}
	}()
	d.Network.Name = d.Name
	err = af.Model(&d.Network).Updates(d.Network).Error
	if err != nil {
		return
	}
	return af.Model(&d).Updates(d).Error
}

func (d Device) UpdateDisabled() (err error) {
	var cnt int
	if dao.DB().Model(&d).Count(&cnt); cnt == 0 {
		return errors.New("invalid ID")
	}
	return dao.DB().Model(&d).Update("disabled", d.Disabled).Error
}

func (d Device) Delete() (err error) {
	db := dao.DB()
	af := db.Begin()
	defer func() {
		if err != nil {
			af.Rollback()
		} else {
			af.Commit()
		}
	}()
	err = af.Unscoped().Delete(deviceNetwork.DeviceNetwork{}, "id = ?", d.NetworkID).Error
	if err != nil {
		return
	}
	return af.Unscoped().Delete(&d).Error
}

func Get(id uint) (o Device, err error) {
	db := dao.DB()
	af := db.Begin()
	defer func() {
		if err != nil {
			af.Rollback()
		} else {
			af.Commit()
		}
	}()
	err = af.Where("id=?", id).First(&o).Error
	if err != nil {
		return
	}
	return o, af.Where("id=?", o.NetworkID).First(&o.Network).Error
}

func (d Device) Get() *Device {
	_d, err := Get(d.ID)
	if err != nil {
		return nil
	}
	return &_d
}

func (d Device) NewPhoto() (img []byte, mark string, err error) {
	bs, err := d.BaseStation()
	if err != nil {
		return
	}
	resp, err := http.Get(
		fmt.Sprintf("http://%v:%v/api/device/%v/newPhoto", bs.IP, 3000, d.Network.IP_alloc),
	)
	if err != nil {
		err = fmt.Errorf("与基站通信时出现错误: %v", err.Error())
		if bs.Online {
			bs.Online = false
			_ = bs.UpdateOnline()
		}
		return
	} else {
		if !bs.Online {
			bs.Online = true
			_ = bs.UpdateOnline()
		}
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	arr := bytes.SplitN(b, []byte{0}, 2)
	if len(arr) != 2 {
		err = errors.New("invalid response format")
		return
	}
	img = arr[1]
	mark = string(arr[0])
	if mark == "!ERROR" {
		err = fmt.Errorf("基站报告了一个错误: %v", string(arr[1]))
	}
	return
}
