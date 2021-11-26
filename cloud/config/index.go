package config

import (
	"github.com/stevenroose/gonfig"
	"log"
	"os"
)

type Config struct {
	DBUsername       string `id:"db_username" default:"root"`
	DBPassword       string `id:"db_password" default:"Talented10"`
	DBHost           string `id:"db_host" default:"localhost"`
	DBPort           int    `id:"db_port" default:"3306"`
	DBDatabase       string `id:"db_database" default:"grip"`
	AssetDir         string `id:"asset_dir" default:"/home/talented/Desktop/grid_presentation/"`
	ListeningAddress string `id:"listening_address" default:"0.0.0.0"`
	ListeningPort    string `id:"listening_port" default:"8110"`
	Secret           string `id:"secret" default:"grip-..."`
}

var config Config

func init() {
	c := gonfig.Conf{
		FileDisable:       true,
		FlagIgnoreUnknown: false,
		EnvPrefix:         "GRIP_",
	}
	err := gonfig.Load(&config, c)
	if config.DBUsername == "" || config.DBPassword == "" {
		log.Fatal("Neither " + c.EnvPrefix + "DB_USERNAME nor " + c.EnvPrefix + "DB_PASSWORD can not be empty! See help by command line flag --help.")
	}
	if err != nil {
		if err.Error() != "unexpected word while parsing flags: '-test.v'" {
			log.Fatal(err)
		}
	}
	_ = os.MkdirAll(config.AssetDir, os.ModeDir|0755)
}

func Get() *Config {
	return &config
}
