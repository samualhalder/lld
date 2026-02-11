package parkinglot

import (
	"fmt"

	"github.com/samualhalder/lld/parkinglot_system/vehicle"
)

type Rate struct {
	Rates map[vehicle.VehicleType]int
}

func (r *Rate) AddRate(vType vehicle.VehicleType, rate int) {
	r.Rates[vType] = rate
}

func (r *Rate) GetRate(vt vehicle.VehicleType) (int, error) {
	rate, ok := r.Rates[vt]
	if !ok {
		return 0, fmt.Errorf("rate not presant")
	}
	return rate, nil
}
