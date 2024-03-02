package contracts

type RegisterDoctor struct {
	Name string `json:"name"`
	Specialization string `json:"specialization"`
	Rating int16 `json:"rating"`
}

type AllSlotsResponse struct {
	Name string `json:"name"`
	Slots []int `json:"slots"`
}