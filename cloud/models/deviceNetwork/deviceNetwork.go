package deviceNetwork

import "github.com/jinzhu/gorm"

type DeviceNetwork struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Auth     string `gorm:"not null"`
	IMSI     string `gorm:"not null"`
	Key      string `gorm:"not null"`
	OP_Type  string `gorm:"not null"`
	OP       string `gorm:"not null"`
	AMF      string `gorm:"not null"`
	SQN      string `gorm:"not null"`
	QCI      string `gorm:"not null"`
	IP_alloc string `gorm:"not null"`
}