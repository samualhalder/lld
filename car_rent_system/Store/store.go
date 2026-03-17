package store

import (
	"github.com/samualhalder/lld/car_rent_system/location"
	"github.com/samualhalder/lld/car_rent_system/vehicles"
)

type Store struct {
	Id       int
	Location *location.Location
	Vehicles []*vehicles.Vehicle
}

func (s *Store) GetAllVehicles() []*vehicles.Vehicle {
	var vehicles []*vehicles.Vehicle
	return vehicles
}

func (s *Store) GetAllAvailable() []*vehicles.Vehicle {
	var vehicles []*vehicles.Vehicle
	return vehicles
}

func (s *Store) GetAllVehiclesByType(vType vehicles.VehicleType) []*vehicles.Vehicle {
	var vehicles []*vehicles.Vehicle
	return vehicles
}
