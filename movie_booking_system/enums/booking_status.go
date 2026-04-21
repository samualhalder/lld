package enums

type BookingStatus int

const (
	Initiated BookingStatus = iota
	Confirmed
	Cancelled
)
