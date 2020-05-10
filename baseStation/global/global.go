package global

import (
	"encoding/csv"
	"github.com/mzz2017/grip/basestation/config"
	"github.com/mzz2017/grip/basestation/models"
	"io/ioutil"
	"log"
	"path"
	"strings"
)

var MapIPDevice map[string]*models.DeviceNetwork

func LoadMapIPDevice(b []byte) {
	var err error
	if b == nil {
		b, err = ioutil.ReadFile(path.Join(config.Get().SrslteConfigDir, "user_db.csv"))
		if err != nil {
			panic(err)
		}
	}
	lines := strings.Split(string(b), "\n")
	t := make([]string, 0, len(lines)+1)
	for _, line := range lines {
		if !strings.HasPrefix(strings.TrimSpace(line), "#") {
			t = append(t, line)
		}
	}
	c := csv.NewReader(strings.NewReader(strings.Join(t, "\n")))
	ds, err := c.ReadAll()
	if err != nil {
		panic(err)
	}
	MapIPDevice = make(map[string]*models.DeviceNetwork)
	for _, line := range ds {
		var d = models.DeviceNetwork{
			Name:     line[0],
			Auth:     line[2],
			IMSI:     line[1],
			Key:      line[3],
			OP_Type:  line[4],
			OP:       line[5],
			AMF:      line[6],
			SQN:      line[7],
			QCI:      line[8],
			IP_alloc: line[9],
		}
		if d.IP_alloc == "dynamic" {
			log.Println("[warning] There is an IP_alloc of a device is 'dynamic'")
		} else if _, ok := MapIPDevice[d.IP_alloc]; ok {
			panic("[error] Duplicated IP_alloc exists: " + d.IP_alloc)
		}
		MapIPDevice[d.IP_alloc] = &d
	}
}

func init() {
	LoadMapIPDevice(nil)
}
