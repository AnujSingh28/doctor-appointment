package main

import (
	"doc-appointment/controller"
	"doc-appointment/service"
	"log"
)

func main(){
	log.Println("Welcome to App")
	var err error

	controller.RegisterDoctor("Dr. Curious", "Cardiologist")
	err = controller.CreateSlotForDoctor("Dr. Curious", "9:30am", "10:00am")
	if err != nil {
		log.Println(err)
	} else {
		log.Println("slot booked")
	}
	log.Println(service.ShowAvailableDoctorsAndSlots("Dr. Curious"))

	err = controller.CreateSlotForDoctor("Dr. Curious", "11:30am", "12:00pm")
	if err != nil {
		log.Println(err)
	} else {
		log.Println("slot booked")
	}
	log.Println(service.ShowAvailableDoctorsAndSlots("Dr. Curious"))

	err = controller.RegisterPatient("anuj")
	if err != nil {
		log.Println(err)
	} else {
		log.Println("patient registered")
	}
	err = controller.RegisterPatient("ankit")
	if err != nil {
		log.Println(err)
	} else {
		log.Println("patient registered")
	}

	booking, err := controller.CreateBookingByDocName("Dr. Curious", "anuj","9:30am", "10:00am")
	if err != nil {
		log.Println(err)
	} else {
		log.Println("booking done: ", booking)
	}
	booking, err = controller.CreateBookingByDocName("Dr. Curious", "ankit","9:30am", "10:00am")
	if err != nil {
		log.Println(err)
	} else {
		log.Println("booking done: ", booking)
	}

	err = service.CancelBookingByPatient("Dr. Curious", "anuj", 1)
	if err != nil {
		log.Println(err)
	}

}
