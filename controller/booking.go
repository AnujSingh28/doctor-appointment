package controller

import (
	"doc-appointment/constants"
	"doc-appointment/models"
	"time"
)

func CreateBookingByDocName(docName, patientName, startTime, endTime string) {

}

func CreateBookingBySpecialization(specialization, patientName, startTime, endTime string) {

}

func CancelBooking() {

}

func BookFromWaitlist() {

}

func GenerateSlot(start time.Time, duration int) models.Slot {
	return models.Slot{Start: start, End: start.Add(time.Minute * constants.Duration)}
}