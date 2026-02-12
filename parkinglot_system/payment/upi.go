package payment

import "github.com/samualhalder/lld/parkinglot_system/ticket"

type Upi struct {
}

func (u *Upi) Pay(ticket *ticket.Ticket) error {
	ticket.PaymentStatus = "Paid"
	return nil
}
