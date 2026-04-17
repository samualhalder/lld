package models

import (
	"fmt"
	"time"

	"github.com/samualhalder/lld/movie_booking_system/utils"
)

type Theater struct {
	Id    string
	Slots map[time.Time][]*Slot
}

func NewTheater(id string) *Theater {
	return &Theater{
		Id:    id,
		Slots: make(map[time.Time][]*Slot),
	}
}

func (t *Theater) AddSlot(slot *Slot, date time.Time) error {
	key := utils.NormalizeDate(date)
	preSlots := t.Slots[key]
	for _, preSlot := range preSlots {
		if preSlot.Screen.Id != slot.Screen.Id {
			continue
		}
		if slot.StartTime.After(preSlot.StartTime) && slot.StartTime.Before(preSlot.EndTime) {
			return fmt.Errorf("overlaping slot")
		}
		if slot.EndTime.After(preSlot.StartTime) && slot.EndTime.Before(preSlot.EndTime) {
			return fmt.Errorf("overlaping slot")
		}
	}
	preSlots = append(preSlots, slot)
	return nil
}

func (t *Theater) GetMovieSlot(movieName string, date time.Time) []*Slot {
	key := utils.NormalizeDate(date)
	var slots []*Slot
	for _, slot := range t.Slots[key] {
		if slot.Movie.Name == movieName {
			slots = append(slots, slot)
		}
	}
	return slots
}
