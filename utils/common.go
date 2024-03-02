package utils

import (
	"time"
)

func ParseTime(timeStamp string) (parsedTime time.Time, err error) {
	layout := "3:04 PM"
	parsedTime, err = time.Parse(layout, timeStamp)
	return
}

func CheckElementExistInSlice(slice []string, element string) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}
