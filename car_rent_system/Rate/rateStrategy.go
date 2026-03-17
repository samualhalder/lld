package rate

import "github.com/samualhalder/lld/car_rent_system/vehicles"

type RateStretagy interface {
	FindRate(vehicle *vehicles.Vehicle) (int, error)
}
