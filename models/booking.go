package models

import "github.com/google/uuid"

type Booking struct {
	DcotorId uuid.UUID `json:"dcotor_id"`
	PatientId uuid.UUID `json:"patient_id"`
	Slot Slot `json:"slot"`
	Waitlist bool `json:"waitlist"`
}
