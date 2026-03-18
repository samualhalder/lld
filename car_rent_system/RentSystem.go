package main

import (
	"fmt"
	"time"

	booking "github.com/samualhalder/lld/car_rent_system/Booking"
	rate "github.com/samualhalder/lld/car_rent_system/Rate"
	store "github.com/samualhalder/lld/car_rent_system/Store"
	"github.com/samualhalder/lld/car_rent_system/location"
	"github.com/samualhalder/lld/car_rent_system/payment"
	"github.com/samualhalder/lld/car_rent_system/vehicles"
)

type RentSystem struct {
	stores   map[*location.Location]*store.Store
	Rate     rate.RateStretagy
	Payment  payment.Payment
	bookings []*booking.Booking
}

func (r *RentSystem) Get(store *store.Store, vehType vehicles.VehicleType) []*vehicles.Vehicle {
	st := r.stores[store.Location]
	return st.GetAllVehiclesByType(vehType)
}

func (r *RentSystem) Book(store *store.Store, car *vehicles.Vehicle) error {
	if car.Status != vehicles.Available {
		return fmt.Errorf("car is not allowed to book")
	}
	booking := &booking.Booking{
		Store:   store,
		Vehicle: car,
		From:    time.Now(),
		To:      time.Now().AddDate(0, 0, 1),
	}
	rent, err := r.Rate.FindRate(booking)
	if err != nil {
		return fmt.Errorf("booking failed")
	}
	err = r.Payment.Pay(rent)
	if err != nil {
		return err
	}
	r.bookings = append(r.bookings, booking)
	return nil
}
