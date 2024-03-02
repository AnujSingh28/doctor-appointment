package contracts

import "doc-appointment/models"

type RegisterDoctor struct {
	Name string `json:"name"`
	Specialization string `json:"specialization"`
	Rating int16 `json:"rating"`
}

type AllSlotsResponse struct {
	Name string `json:"name"`
	Slots []models.Slot `json:"slots"`
}