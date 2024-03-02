package controller

import (
	"doc-appointment/constants"
	"doc-appointment/contracts"
	"doc-appointment/models"
	"doc-appointment/service"
	"doc-appointment/utils"
	"errors"
	"github.com/google/uuid"
)

func RegisterDoctor (name, specialization string, rating int) (err error) {
	docExists := service.IsDoctorAlreadyExists(name)
	if docExists{
		err = errors.New("Doctor already exists")
		return
	}

	newDoctor := models.Doctor{
		Id:             uuid.New(),
		Name:           name,
		Specialization: specialization,
		Rating:         rating,
	}
	err = service.RegisterDoctor(newDoctor)
	return err
}

func SearchAllDoctorSlotsBasedOnSpeciality(speciality string) ([]contracts.AllSlotsResponse, error) {

	if utils.CheckElementExistInSlice(constants.DoctorSpecialities, speciality) {
		return service.SearchAllDoctorSlotsBasedOnSpeciality(speciality)
	}
	return []contracts.AllSlotsResponse{}, nil
}

func CreateSlotForDoctor(docName, start, end string) error {
	if !service.IsDoctorAlreadyExists(docName) {
		// need specialization of doctor here if want to register
		return errors.New("Doctor doesn't exists!")
	}
	startTime, err := utils.ParseTime(start)
	if err != nil {
		return err
	}
	endTime, err := utils.ParseTime(start)
	if err != nil {
		return err
	}

	err = service.CreateSlotForDoctor(docName, startTime, endTime)
	return err
}

func AllBookedAppointmentsForDoctor (docName string) ([]models.Slot, error) {
	if !service.IsDoctorAlreadyExists(docName) {
		// need specialization of doctor here if want to register
		return []models.Slot{}, errors.New("Doctor doesn't exists!")
	}
	return service.SearchAllBookedAppointmentsForDoctor(docName)
}
