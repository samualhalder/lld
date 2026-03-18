package rate

import (
	"time"

	booking "github.com/samualhalder/lld/car_rent_system/Booking"
)

type DefaultStrategy struct {
}

func (d *DefaultStrategy) FindRate(booking *booking.Booking) (int, error) {
	diff := booking.To.Sub(booking.From)
	days := int(diff / (24 * time.Hour))
	return days * booking.Vehicle.BasePrice, nil
}
