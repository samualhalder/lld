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
	Tickets map[int]*ticket.Ticket
}

func (f *Floor) AddSpace(vt vehicle.VehicleType, cnt int) {
	f.Space[vt] = cnt
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
		f.Tickets[v.GetId()] = tkt
		return tkt, nil
	}
	return nil, fmt.Errorf("no space available")
}

func (f *Floor) UnPark(v vehicle.Vehicle) (*ticket.Ticket, bool) {
	tkt, ok := f.Tickets[v.GetId()]
	if !ok {
		return nil, false
	}
	tkt.ExitTime = time.Now()
	f.Space[v.IsType()]++
	return tkt, true
}
