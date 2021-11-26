package dao

import (
	"github.com/jinzhu/gorm"
)
import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mzz2017/grip/cloud/config"
	"sync"
)

var once sync.Once
var db *gorm.DB

func initDB() {
	conf := config.Get()
	var args = fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local",
		conf.DBUsername, conf.DBPassword, conf.DBHost, conf.DBPort, conf.DBDatabase,
	)
	var err error
	db, err = gorm.Open("mysql", args)
	if err != nil {
		panic(err)
	}
}

func DB() *gorm.DB {
	once.Do(initDB)
	return db
}
