package service

import (
	"doc-appointment/contracts"
	"doc-appointment/models"
	"errors"
	"log"
)

var (
	AllDoctors = map[string]models.Doctor{} // name _> model.doctor
	AvailableDoctorsAndSlots = map[string]models.AvailableDoctor{} // drname -> {name, slots []int}
	AllAvailableSlots        = map[int]bool{}
)

func IsDoctorAlreadyExists(name string) bool {
	if _, ok := AllDoctors[name]; ok {
		return true
	}
	return false
}

func ShowAvailableDoctorsAndSlots(name string) map[string]models.AvailableDoctor {
	return AvailableDoctorsAndSlots
}

func IsSLotAlreadyCreated (slot int) bool {
	if booked, ok := AllAvailableSlots[slot]; ok {
		return booked
	}
	return false
}

func RegisterDoctor(newDoc models.Doctor) error {
	AllDoctors[newDoc.Name] = newDoc
	return nil
}

func CreateSlotForDoctor(docName string, slot int) error {

	if IsSLotAlreadyCreated(slot) {
		return errors.New("Slot already created by a doctor ")
	}

	if _, ok := AvailableDoctorsAndSlots[docName]; ok {
		// check if slot already exist else push
		log.Println("-----------------")
		reqDoctorStruct := AvailableDoctorsAndSlots[docName]
		docSlots := reqDoctorStruct.Slots
		docSlots = append(docSlots, slot)
		reqDoctorStruct.Slots = docSlots
		AvailableDoctorsAndSlots[docName] = reqDoctorStruct
	} else {
		newSlotForDoc := models.AvailableDoctor{
			DocName: docName,
			Slots: []int{slot},
		}
		AvailableDoctorsAndSlots[docName] = newSlotForDoc
	}


	AllAvailableSlots[slot] = true
	return nil
}


func SearchAllDoctorSlotsBasedOnSpeciality(speciality string) ([]contracts.AllSlotsResponse, error) {

	var allSlotsAvailable []contracts.AllSlotsResponse
	for docName, availability := range AvailableDoctorsAndSlots {
		//doctor, err := GetDoctorById(availability.DoctorId)
		//if err != nil {
		//	return allSlotsAvailable, err
		//}
		if AllDoctors[docName].Specialization == speciality {
			allSlotsAvailable = append(allSlotsAvailable, contracts.AllSlotsResponse{
				Name:  docName,
				Slots: availability.Slots,
			})
		}
	}
	return allSlotsAvailable, nil
}

func SearchAllBookedAppointmentsForDoctor(docName string) ([]int, error) {
	return AvailableDoctorsAndSlots[docName].Slots, nil
}