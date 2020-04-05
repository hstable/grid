package config

import (
	"github.com/gin-gonic/gin"
	"github.com/stevenroose/gonfig"
	"log"
)

type Config struct {
	DBFile           string `id:"db_file" default:"grip_basestation.db"`
	ListeningAddress string `id:"listening_address" default:"0.0.0.0"`
	ListeningPort    string `id:"listening_port" default:"3000"`
	Secret           string `id:"secret" default:"grip-..."`
	SrslteConfigDir  string `id:"srslte_config_dir" default:"~/.config/srslte"`
	ServerAddress    string `id:"server_address" default:"https://grip.mzz.pub:8110"`
}

var config Config

func init() {
	c := gonfig.Conf{
		FileDisable:       true,
		FlagIgnoreUnknown: false,
		EnvPrefix:         "GRIP_",
	}
	err := gonfig.Load(&config, c)
	if err != nil {
		if err.Error() != "unexpected word while parsing flags: '-test.v'" {
			log.Fatal(err)
		}
	}
	if gin.IsDebugging() {
		config.ServerAddress = "http://localhost:8110"
	}
}

func Get() *Config {
	return &config
}
