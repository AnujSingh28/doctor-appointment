package utils

import (
	"fmt"
	"time"
)

func ParseTime(timeStamp string) (parsedTime time.Time, err error) {
	layout := "3:04 PM"
	parsedTime, err = time.Parse(layout, timeStamp)
	return
}

func CheckElementExistInSlice(slice []int, element int) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}

func GetTimeSlot(startTime, endTime string) (int, error) {
	timeFormat := "3:04pm"
	start, err := time.Parse(timeFormat, startTime)
	if err != nil {
		return 0, fmt.Errorf("error parsing start time: %v", err)
	}
	end, err := time.Parse(timeFormat, endTime)
	if err != nil {
		return 0, fmt.Errorf("error parsing end time: %v", err)
	}

	startRange, _ := time.Parse(timeFormat, "9:00am")
	endRange, _ := time.Parse(timeFormat, "9:00pm")

	if start.Before(startRange) || end.After(endRange) {
		return 0, fmt.Errorf("time should be in range 9am to 9pm")
	}

	slotDuration := 30*time.Minute
	currentTime := start
	slotCount := 0

	for currentTime.Before(end) {
		if currentTime.After(startRange) && currentTime.Before(endRange) {
			slotCount++
		}
		currentTime = currentTime.Add(slotDuration)

	}
	return slotCount, nil
}

func GetStartAndEndFromSlot (slotNumber int) (string, string, error) {
	timeFormat := "3:04pm"
	startRange, _ := time.Parse(timeFormat, "9:00am")
	slotDuration := 30*time.Minute

	startTime := startRange.Add(time.Duration(slotNumber)*slotDuration)

	if startTime.Before(startRange) || startTime.After(startRange.Add(12*time.Hour)) {
		return "", "", fmt.Errorf("invalid slot number")
	}

	endTime := startTime.Add(slotDuration)
	return startTime.Format(timeFormat), endTime.Format(timeFormat), nil

}