package main

import (
	"strings"
	"time"

	"github.com/samualhalder/lld/movie_booking_system/enums"
	"github.com/samualhalder/lld/movie_booking_system/models"
	paymentstrategy "github.com/samualhalder/lld/movie_booking_system/payment_strategy"
)

type BookingController struct {
	theaters []*models.Theater
}

func (b *BookingController) FindTheaters(movieName string) map[*models.Theater][]*models.Slot {
	movieName = strings.ToLower(movieName)
	mp := make(map[*models.Theater][]*models.Slot)
	today := time.Now()
	for _, theater := range b.theaters {

		for time, slots := range theater.Slots {
			if time.Before(today) {
				continue
			}
			for _, slot := range slots {
				if strings.ToLower(slot.Movie.Name) == movieName {
					mp[theater] = append(mp[theater], slot)
				}
			}
		}
	}
	return mp
}

func (b *BookingController) SelectSeat(theater *models.Theater, slot *models.Slot, seats []*models.Seat, user *models.User) *models.Booking {
	booking, _ := theater.SelectSeat(user, slot, seats)

	return booking
}

func (b *BookingController) Pay(paymentStrategy *paymentstrategy.Upi, booking *models.Booking) {
	paymentStrategy.Pay(booking.Total)
	booking.Status = enums.Confirmed
	booking.Paid = true
	booking.Theater.ConfirmSeat(booking.User, booking.Slot, booking.Seats)
}
