package payment

import "github.com/samualhalder/lld/parkinglot_system/ticket"

type Payment interface {
	Pay(ticket *ticket.Ticket) error
}
