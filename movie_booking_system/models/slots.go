package models

import (
	"time"

	"github.com/samualhalder/lld/movie_booking_system/lock"
)

type Slot struct {
	Id              string
	Screen          *Screen
	Date            time.Time
	StartTime       time.Time
	EndTime         time.Time
	Seats           map[string]*Seat
	Movie           *Movie
	LockingStrategy lock.Lock
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

func (s *Slot) LockSeat(seat *Seat, user *User) error {
	err := s.LockingStrategy.SetLock(user.Id, s.Id, seat.Id)
	if err != nil {
		return err
	}
	return nil
}
