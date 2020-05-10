package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"strings"
)

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

type DeviceNetworks []DeviceNetwork

func (dn DeviceNetworks) ToCSV() string {
	var builder strings.Builder
	builder.WriteString(
		`#
# .csv to store UE's information in HSS
# Kept in the following format: "Name,Auth,IMSI,Key,OP_Type,OP,AMF,SQN,QCI,IP_alloc"
#
# Name:     Human readable name to help distinguish UE's. Ignored by the HSS
# Auth:     Authentication algorithm used by the UE. Valid algorithms are XOR
#           (xor) and MILENAGE (mil)
# IMSI:     UE's IMSI value
# Key:      UE's key, where other keys are derived from. Stored in hexadecimal
# OP_Type:  Operator's code type, either OP or OPc
# OP/OPc:   Operator Code/Cyphered Operator Code, stored in hexadecimal
# AMF:      Authentication management field, stored in hexadecimal
# SQN:      UE's Sequence number for freshness of the authentication
# QCI:      QoS Class Identifier for the UE's default bearer.
# IP_alloc: IP allocation stratagy for the SPGW.
#           With 'dynamic' the SPGW will automatically allocate IPs
#           With a valid IPv4 (e.g. '172.16.0.2') the UE will have a statically assigned IP.
#
# Note: Lines starting by '#' are ignored and will be overwritten
`)
	for _, d := range dn {
		builder.WriteString(fmt.Sprintf("%v,%v,%v,%v,%v,%v,%v,%v,%v,%v\n", d.Name, d.Auth, d.IMSI, d.Key, d.OP_Type, d.OP, d.AMF, d.SQN, d.QCI, d.IP_alloc))
	}
	return builder.String()
}
