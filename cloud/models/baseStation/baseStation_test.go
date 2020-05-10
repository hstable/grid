package baseStation

import (
	"github.com/jinzhu/gorm"
	"testing"
)

func TestBaseStation_PushDevices(t *testing.T) {
	bs := BaseStation{
		Model: gorm.Model{
			ID: 1,
		},
		Name:     "test",
		IP:       "127.0.0.1",
		Location: "0.0,0.0",
		Online:   true,
	}
	err := bs.PushDevices()
	if err != nil {
		t.Fatal(err)
	}
}
