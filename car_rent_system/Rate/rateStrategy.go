package rate

import (
	booking "github.com/samualhalder/lld/car_rent_system/Booking"
)

type RateStretagy interface {
	FindRate(booking *booking.Booking) (int, error)
}
