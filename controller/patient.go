package controller

import (
	"doc-appointment/models"
	"doc-appointment/service"
	"errors"
)

func RegisterPatient (name string) (err error ){
	docExists := service.IsPatientAlreadyExists(name)
	if docExists{
		err = errors.New("Patient already exists")
		return
	}

	newPatient := models.Patient{
		Name:           name,
	}
	err = service.RegisterPatient(newPatient)
	return err
}

func AllBookedAppointmentsForPatient () {

}