package service

import (
	"doc-appointment/models"
	"log"
)

var (
	AllPatients = map[string]models.Patient{}
)

func IsPatientAlreadyExists(name string) bool {
	if _, ok := AllPatients[name]; ok {
		return true
	}
	return false
}

func RegisterPatient(newPatient models.Patient) error {
	AllPatients[newPatient.Name] = newPatient
	log.Println(AllPatients)
	return nil
}

