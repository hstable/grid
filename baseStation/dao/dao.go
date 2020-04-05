package dao

import (
	"github.com/jinzhu/gorm"
)
import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/mzz2017/grip/basestation/config"
	"sync"
)

var once sync.Once
var db *gorm.DB

func initDB() {
	conf := config.Get()
	var err error
	db, err = gorm.Open("sqlite3", conf.DBFile)
	if err != nil {
		panic(err)
	}
}

func DB() *gorm.DB {
	once.Do(initDB)
	return db
}
