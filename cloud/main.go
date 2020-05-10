package main

import (
	"fmt"
	"github.com/mzz2017/grip/cloud/config"
	"github.com/mzz2017/grip/cloud/dao"
	"github.com/mzz2017/grip/cloud/models/baseStation"
	"github.com/mzz2017/grip/cloud/models/company"
	"github.com/mzz2017/grip/cloud/models/device"
	"github.com/mzz2017/grip/cloud/models/deviceNetwork"
	"github.com/mzz2017/grip/cloud/models/image"
	"github.com/mzz2017/grip/cloud/models/line"
	"github.com/mzz2017/grip/cloud/models/power"
	"github.com/mzz2017/grip/cloud/models/tower"
	"github.com/mzz2017/grip/cloud/models/user"
	"github.com/mzz2017/grip/cloud/router"
	"log"
)

func initDB() {
	db := dao.DB()
	//自动迁移Tables
	db.Set("gorm:table_options", "ENGINE=innodb,DEFAULT CHARSET=utf8mb4").AutoMigrate(
		&company.Company{},
		&image.Image{},
		&user.User{},
		&baseStation.BaseStation{},
		&deviceNetwork.DeviceNetwork{},
		&device.Device{},
		&line.Line{},
		&power.Power{},
		&tower.Tower{},
	)
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	initDB()
	conf := config.Get()
	address := conf.ListeningAddress
	port := conf.ListeningPort
	log.Fatal(router.Router.Run(fmt.Sprintf("%s:%s", address, port)))
}
