package rate

import (
	"time"

	"github.com/samualhalder/lld/car_rent_system/vehicles"
)

type WeekendStrategy struct {
	PriceHike int
}

func (d *WeekendStrategy) FindRate(vehicle *vehicles.Vehicle) (int, error) {
	day := time.Now().Weekday()
	if day == 0 || day == 6 {
		return vehicle.BasePrice * (100 + d.PriceHike) / 100, nil
	}
	return vehicle.BasePrice, nil
}
