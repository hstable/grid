package location

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Location string

func New(lat float64, lon float64) Location {
	return Location(fmt.Sprintf("%v,%v", lat, lon))
}
func (l Location) Get() (lat, lon float64) {
	arr := strings.SplitN(string(l), ",", 2)
	lat, _ = strconv.ParseFloat(arr[0], 64)
	lon, _ = strconv.ParseFloat(arr[1], 64)
	return
}
func Parse(location string) (l Location, err error) {
	arr := strings.Split(string(location), ",")
	if len(arr) != 2 {
		err = errors.New("invalid format")
		return
	}
	_, err1 := strconv.ParseFloat(arr[0], 64)
	_, err2 := strconv.ParseFloat(arr[1], 64)
	if err1 != nil || err2 != nil {
		err = errors.New("invalid format")
		return
	}
	return Location(location), nil
}
