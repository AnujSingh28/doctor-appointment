package models

var Specialization  = map[string]string {
	"Cardiologist": "cardiologist",
	"Dermatologist": "dermatologist",
	"Orthopedic": "dermatologist",
	"General Physician": "general_physician",
}

type Doctor struct {
	Name string `json:"name"`
	Specialization string `json:"specialization"`
	Rating int `json:"rating"`
}

type AvailableDoctor struct {
	DocName string `json:"doc_name"`
	Slots []int `json:"slots"`
}

