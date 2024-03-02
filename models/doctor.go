package models

import (
	"github.com/google/uuid"
	"time"
)

var Specialization  = map[string]string {
	"Cardiologist": "cardiologist",
	"Dermatologist": "dermatologist",
	"Orthopedic": "dermatologist",
	"General Physician": "general_physician",
}

type Doctor struct {
	Id uuid.UUID `json:"id"`
	Name string `json:"name"`
	Specialization string `json:"specialization"`
	Rating int `json:"rating"`
}

type Slot struct {
	Start time.Time
	End time.Time
}

type AvailableDoctor struct {
	DoctorId uuid.UUID `json:"doctor_id"`
	Slots []Slot `json:"slots"`
}

