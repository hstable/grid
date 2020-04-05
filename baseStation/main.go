package main

import (
	"fmt"
	"github.com/mzz2017/grip/basestation/config"
	"github.com/mzz2017/grip/basestation/dao"
	"github.com/mzz2017/grip/basestation/router"
	"log"
)

func initDB() {
	db := dao.DB()
	//自动迁移Tables
	db.AutoMigrate(

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
