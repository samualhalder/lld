package parkinglot

import (
	"fmt"
	"time"

	"github.com/samualhalder/lld/parkinglot_system/ticket"
	"github.com/samualhalder/lld/parkinglot_system/vehicle"
)

type Floor struct {
	Id      int
	Space   map[vehicle.VehicleType]int
	tickets map[int]*ticket.Ticket
}

func (f *Floor) Park(v vehicle.Vehicle) (*ticket.Ticket, error) {
	val, ok := f.Space[v.IsType()]
	if !ok {
		return nil, fmt.Errorf("no space available")
	}
	if val > 0 {
		f.Space[v.IsType()]--
		tkt := &ticket.Ticket{
			VehidleId: v.GetId(),
			EntryTime: time.Now(),
		}
		f.tickets[v.GetId()] = tkt
		return tkt, nil
	}
	return nil, fmt.Errorf("no space available")
}

func (f *Floor) UnPark(v vehicle.Vehicle) (*ticket.Ticket, bool) {
	tkt, ok := f.tickets[v.GetId()]
	if !ok {
		return nil, false
	}
	tkt.EntryTime = time.Now()
	return tkt, true
}
