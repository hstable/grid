package global

import (
	"encoding/csv"
	"github.com/mzz2017/grip/basestation/config"
	"github.com/mzz2017/grip/basestation/models"
	"os"
	"path"
)

var MapIPDevice map[string]*models.Device

func init() {
	f, err := os.Open(path.Join(config.Get().SrslteConfigDir, "user_db.csv"))
	if err != nil {
		panic(err)
	}
	c := csv.NewReader(f)
	ds, err := c.ReadAll()
	if err != nil {
		panic(err)
	}
	MapIPDevice = make(map[string]*models.Device)
	for _, line := range ds {
		var d = models.Device{
			Name:     line[0],
			IMSI:     line[1],
			Auth:     line[2],
			Key:      line[3],
			OPType:   line[4],
			OPc:      line[5],
			AMF:      line[6],
			SQN:      line[7],
			QCI:      line[8],
			IP_alloc: line[9],
		}
		MapIPDevice[d.IP_alloc] = &d
	}
}
