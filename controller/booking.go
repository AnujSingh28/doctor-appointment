package controller

import (
	"doc-appointment/models"
	"doc-appointment/service"
	"doc-appointment/utils"
)

func CreateBookingByDocName(docName, patientName, startTime, endTime string) (models.Booking, error){
	reqSlot, err := utils.GetTimeSlot(startTime, endTime)
	if err != nil {
		return models.Booking{}, err
	}
	return service.CreateBookingByDocName(docName, patientName, reqSlot)

}

//func CreateBookingBySpecialization(specialization, patientName, startTime, endTime string) (models.Booking, error) {
//	reqSlot, err := utils.GetTimeSlot(startTime, endTime)
//	if err != nil {
//		return models.Booking{}, err
//	}
//	return service.CreateBookingBySpecialization(specialization, patientName, reqSlot)
//}

//func CancelBooking() {
//
//}