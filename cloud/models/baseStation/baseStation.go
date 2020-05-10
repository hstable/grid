package baseStation

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/mzz2017/grip/cloud/common"
	"github.com/mzz2017/grip/cloud/dao"
	"github.com/mzz2017/grip/cloud/dataStructure/location"
	"github.com/mzz2017/grip/cloud/models/deviceNetwork"
	"io/ioutil"
	"net/http"
)

type BaseStation struct {
	gorm.Model
	Name     string            `gorm:"not null;unique"`
	IP       string            `gorm:"not null"`
	Location location.Location `gorm:"not null"`
	Online   bool              `gorm:"not null"`
}

func (b BaseStation) Insert() error {
	return dao.DB().Model(&BaseStation{}).Create(&b).Error
}

func (b BaseStation) Update() (err error) {
	var e BaseStation
	if dao.DB().Model(&b).First(&e); e.CreatedAt.IsZero() {
		return errors.New("invalid ID")
	}
	b.Online = e.Online
	return dao.DB().Model(&b).Updates(b).Error
}

func (b BaseStation) UpdateOnline() (err error) {
	var cnt int
	if dao.DB().Model(&b).Count(&cnt); cnt == 0 {
		return errors.New("invalid ID")
	}
	return dao.DB().Model(&b).Update("online", b.Online).Error
}

func Get(id uint) (o BaseStation, err error) {
	db := dao.DB()
	err = db.Where("id=?", id).First(&o).Error
	return
}
func (b BaseStation) Delete() error {
	return dao.DB().Unscoped().Delete(&b).Error
}

func getDeviceNetworks(baseStationID uint) (dn []*deviceNetwork.DeviceNetwork, err error) {
	dn = make([]*deviceNetwork.DeviceNetwork, 0)
	rows, err := dao.DB().Raw("select device_networks.* from device_networks,devices,towers where base_station_id=? and tower_id=towers.id and network_id=device_networks.id", baseStationID).Rows()
	if err != nil {
		return
	}
	for rows.Next() {
		n := new(deviceNetwork.DeviceNetwork)
		err := dao.DB().ScanRows(rows, n)
		if err != nil {
			continue
		}
		dn = append(dn, n)
	}
	return
}
func (b BaseStation) PushDevices() (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("PushDevices: %v", err.Error())
		}
	}()
	devices, err := getDeviceNetworks(b.ID)
	if err != nil {
		return
	}
	c, err := json.Marshal(gin.H{
		"devices": devices,
	})
	if err != nil {
		return
	}
	resp, err := http.Post(
		fmt.Sprintf("http://%v:%v/api/devices", b.IP, 3000),
		"application/json",
		bytes.NewReader(c),
	)
	if err != nil {
		err = fmt.Errorf("与基站通信时出现错误: %v", err.Error())
		return
	}
	defer resp.Body.Close()
	bb, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var res common.ResponseType
	_ = json.Unmarshal(bb, &res)
	if res.Code == common.FAIL {
		return errors.New(*res.Message)
	}
	return
}

func (b BaseStation) Reboot() (err error) {
	//TODO
	//resp, err := http.Post(
	//	fmt.Sprintf("http://%v:%v/api/reboot", b.IP, 3000),
	//	"application/json",
	//	nil,
	//)
	return
}