package service

import (
	"doc-appointment/contracts"
	"doc-appointment/models"
	"errors"
	"github.com/google/uuid"
	"log"
	"time"
)

var (
	AllDoctors = map[string]models.Doctor{}
	AvailableDoctorsAndSlots = map[string]models.AvailableDoctor{}
)

func IsDoctorAlreadyExists(name string) bool {
	if _, ok := AllDoctors[name]; ok {
		return true
	}
	return false
}

func RegisterDoctor(newDoc models.Doctor) error {
	AllDoctors[newDoc.Name] = newDoc
	log.Println(AllDoctors)
	return nil
}

func CreateSlotForDoctor(docName string, start, end time.Time) error {

	if _, ok := AvailableDoctorsAndSlots[docName]; ok {
		// check if slot already exist else push
		reqDoctorStruct := AvailableDoctorsAndSlots[docName]
		reqDoctorStruct.Slots = append(reqDoctorStruct.Slots, models.Slot{
			Start: start,
			End:   end,
		})

	}
	newSlotForDoc := models.AvailableDoctor{
		DoctorId: AllDoctors[docName].Id,
		Slots: []models.Slot{{
			Start: start,
			End:   end,
		}},
	}
	AvailableDoctorsAndSlots[docName] = newSlotForDoc
	return nil
}

func GetDoctorById (id uuid.UUID) (models.Doctor, error) {
	var doctor models.Doctor
	for _, doctor = range AllDoctors {
		if doctor.Id == id {
			return doctor, nil
		}
	}
	return doctor, errors.New("Doctor doesn't exists")
}

func SearchAllDoctorSlotsBasedOnSpeciality(speciality string) ([]contracts.AllSlotsResponse, error) {

	var allSlotsAvailable []contracts.AllSlotsResponse
	for docName, availability := range AvailableDoctorsAndSlots {
		doctor, err := GetDoctorById(availability.DoctorId)
		if err != nil {
			return allSlotsAvailable, err
		}
		if doctor.Specialization == speciality {
			allSlotsAvailable = append(allSlotsAvailable, contracts.AllSlotsResponse{
				Name:  docName,
				Slots: availability.Slots,
			})
		}
	}
	return allSlotsAvailable, nil
}

func SearchAllBookedAppointmentsForDoctor(docName string) ([]models.Slot, error) {
	return AvailableDoctorsAndSlots[docName].Slots, nil
}