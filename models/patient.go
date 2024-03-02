package models

import "github.com/google/uuid"

type Patient struct {
	Id uuid.UUID `json:"id"`
	Name string `json:"name"`
	Appointment PatientAppointment `json:"appointment"`
}

type PatientAppointment struct {
	DoctorId uuid.UUID `json:"doctor_id"`
	Slots []Slot `json:"slots"`
}