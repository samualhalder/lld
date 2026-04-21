package models

import (
	"github.com/samualhalder/lld/movie_booking_system/enums"
	paymentstrategy "github.com/samualhalder/lld/movie_booking_system/payment_strategy"
)

type Booking struct {
	Theater         *Theater
	Slot            *Slot
	Seats           []*Seat
	User            *User
	Total           int
	Paid            bool
	PaymentStrategy paymentstrategy.PaymentStrategy
	Status          enums.BookingStatus
}
