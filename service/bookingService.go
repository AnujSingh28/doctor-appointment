package service

import (
	"doc-appointment/models"
	"doc-appointment/utils"
	"errors"
	"log"
	"sort"
)

var (
	WaitingQueue = map[int][]models.Booking{}
	BookedSlots = map[int]models.Booking{}
)

func IsSlotAlreadyBooked(slot int) bool {
	if _, ok := BookedSlots[slot]; ok {
		return true
	}
	return false
}

func CreateBookingByDocName(docName, patientName string, slot int) (models.Booking, error) {
	if !IsSLotAlreadyCreated(slot) {
		errors.New("This slot is not made available by any doctor ")
	}

	newBooking := models.Booking{
		DoctorName:  docName,
		PatientName: patientName,
		Slot:        slot,
		Waitlist:    false,
	}

	for name, docSlot := range AvailableDoctorsAndSlots {
		if name == docName && utils.CheckElementExistInSlice(docSlot.Slots, slot) {
			if IsSlotAlreadyBooked(slot) {
				// push in waitlist
				newBooking.Waitlist = true
				WaitingQueue[slot] = append(WaitingQueue[slot], newBooking)
				return newBooking, errors.New("Booking in wait list ")
			} else {
				BookedSlots[slot] = newBooking
				return newBooking, nil
			}
		}
	}
	return models.Booking{}, errors.New("no doctor slots available ")
}

func CreateBookingBySpecialization (speciality, patientName string, slot int) (models.Booking, error) {
	docs, err := SearchAllDoctorSlotsBasedOnSpeciality(speciality)
	if err != nil {
		return models.Booking{}, err
	}

	var allSlots []int

	for _, doc := range docs {
		allSlots = append(allSlots, doc.Slots...)
	}

	sort.Slice(allSlots, func(i, j int) bool {
		return allSlots[i] < allSlots[j]
	})

	return models.Booking{}, nil
}

func CancelBookingByPatient(docName, patientName string, slot int) error {
	if !IsSlotAlreadyBooked(slot) {
		return errors.New("This slot is not booked, can't cancel ")
	}

	for bookedSlot, booking := range BookedSlots {
		if slot == bookedSlot && booking.DoctorName == docName && booking.PatientName == patientName {
			delete(BookedSlots, slot)
			log.Println("cancelled booking: ", booking)
			if len(WaitingQueue[slot]) > 0 {
				newBooking := WaitingQueue[slot][0]
				WaitingQueue[slot] = WaitingQueue[slot][1:]
				newBooking.Waitlist = false
				BookedSlots[slot] = newBooking
				log.Println("New booking from waitlist: ", newBooking)
			}
		}
	}
	return nil
}