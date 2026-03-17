package main

import (
	booking "github.com/samualhalder/lld/car_rent_system/Booking"
	rate "github.com/samualhalder/lld/car_rent_system/Rate"
	store "github.com/samualhalder/lld/car_rent_system/Store"
	"github.com/samualhalder/lld/car_rent_system/location"
	"github.com/samualhalder/lld/car_rent_system/payment"
)

type RentSystem struct {
	stores   map[*location.Location]*store.Store
	Rate     rate.RateStretagy
	Payment  payment.Payment
	bookings []*booking.Booking
}
