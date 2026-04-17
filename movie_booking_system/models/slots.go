package models

import "time"

type Slot struct {
	Id        string
	Screen    *Screen
	Date      time.Time
	StartTime time.Time
	EndTime   time.Time
	Seats     map[string]*Seat
	Movie     *Movie
}

func NewSlot(id string) *Slot {
	return &Slot{
		Id:    id,
		Seats: make(map[string]*Seat),
	}
}

func (s *Slot) AddSeat(seat *Seat) {
	s.Seats[seat.Id] = seat
}
