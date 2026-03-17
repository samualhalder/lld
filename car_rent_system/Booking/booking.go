package booking

import (
	"time"

	store "github.com/samualhalder/lld/car_rent_system/Store"
	"github.com/samualhalder/lld/car_rent_system/vehicles"
)

type Booking struct {
	Vehicle     *vehicles.Vehicle
	Store       *store.Store
	From        time.Time
	To          time.Time
	TotalAmount int
}
