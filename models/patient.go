package models

type Patient struct {
	Name string `json:"name"`
	Appointment PatientAppointment `json:"appointment"`
}

type PatientAppointment struct {
	DocName string `json:"doc_name"`
	Slots []int `json:"slots"`
}