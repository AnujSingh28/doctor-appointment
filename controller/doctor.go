package controller

import (
	"doc-appointment/constants"
	"doc-appointment/contracts"
	"doc-appointment/models"
	"doc-appointment/service"
	"doc-appointment/utils"
	"errors"
	"log"
)

func RegisterDoctor (name, specialization string) (newDoctor models.Doctor, err error) {
	docExists := service.IsDoctorAlreadyExists(name)
	if docExists{
		err = errors.New("Doctor already exists")
		return
	}

	newDoctor = models.Doctor{
		Name:           name,
		Specialization: specialization,
		Rating:         constants.DeafultRating,
	}
	err = service.RegisterDoctor(newDoctor)
	if err != nil {
		log.Println(err)
	}
	log.Println("welcome: ", newDoctor.Name)
	return newDoctor, err
}

func SearchAllDoctorSlotsBasedOnSpeciality(speciality string) ([]contracts.AllSlotsResponse, error) {

	return service.SearchAllDoctorSlotsBasedOnSpeciality(speciality)
}

func CreateSlotForDoctor(docName, start, end string) error {
	if !service.IsDoctorAlreadyExists(docName) {
		// need specialization of doctor here if want to register
		return errors.New("Doctor doesn't exists!")
	}

	slot, err := utils.GetTimeSlot(start, end)
	log.Println("-----", slot)
	if err != nil {
		return err
	}

	err = service.CreateSlotForDoctor(docName, slot)
	return err
}

func AllBookedAppointmentsForDoctor (docName string) ([]int, error) {
	if !service.IsDoctorAlreadyExists(docName) {
		// need specialization of doctor here if want to register
		return []int{}, errors.New("Doctor doesn't exists!")
	}
	return service.SearchAllBookedAppointmentsForDoctor(docName)
}
