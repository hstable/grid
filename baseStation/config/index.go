package config

import (
	"github.com/gin-gonic/gin"
	"github.com/stevenroose/gonfig"
	"log"
	"os"
	"strings"
)

type Config struct {
	DBFile           string `id:"db_file" default:"grip_basestation.db"`
	ListeningAddress string `id:"listening_address" default:"0.0.0.0"`
	ListeningPort    string `id:"listening_port" default:"3000"`
	Secret           string `id:"secret" default:"grip-..."`
	SrslteConfigDir  string `id:"srslte_config_dir" default:"~/.config/srslte/"`
	ServerAddress    string `id:"server_address" default:"https://grip.mzz.pub:8110"`
	AssetDir         string `id:"asset_dir" default:"/home/talented/Desktop/grid_presentation/"`
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
	if strings.HasPrefix(config.SrslteConfigDir, "~/") {
		h, _ := os.UserHomeDir()
		config.SrslteConfigDir = strings.Replace(config.SrslteConfigDir, "~", h, 1)
	}

}

func Get() *Config {
	return &config
}
