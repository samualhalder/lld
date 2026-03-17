package rate

import "github.com/samualhalder/lld/car_rent_system/vehicles"

type DefaultStrategy struct {
}

func (d *DefaultStrategy) FindRate(vehicle *vehicles.Vehicle) (int, error) {
	return vehicle.BasePrice, nil
}
