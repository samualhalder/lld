package models

import (
	"github.com/samualhalder/lld/movie_booking_system/enums"
)

type Seat struct {
	Id       string
	Category enums.SeatType
	Price    int
	Booked   bool
}
