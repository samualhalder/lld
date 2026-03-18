package rate

import (
	"time"

	booking "github.com/samualhalder/lld/car_rent_system/Booking"
)

type WeekendStrategy struct {
	PriceHike int
}

func (d *WeekendStrategy) FindRate(booking *booking.Booking) (int, error) {
	start := booking.From.Weekday()
	diff := booking.To.Sub(booking.From)
	days := int(diff / (24 * time.Hour))
	if start == 0 || start == 6 {
		return days * booking.Vehicle.BasePrice * d.PriceHike, nil
	}

	return days * booking.Vehicle.BasePrice , nil
}
