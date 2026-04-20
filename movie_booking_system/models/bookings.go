package models

import paymentstrategy "github.com/samualhalder/lld/movie_booking_system/payment_strategy"

type Booking struct {
	Id              string
	Slot            *Slot
	Seats           []*Seat
	User            *User
	Total           int
	Paid            bool
	PaymentStrategy paymentstrategy.PaymentStrategy
}

