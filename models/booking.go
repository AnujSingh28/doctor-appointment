package models

type Booking struct{
	DoctorName string `json:"doc_name"`
	PatientName string `json:"patient_name"`
	Slot int `json:"slot"`
	Waitlist bool `json:"waitlist"`
}
