package parkinglot

import (
	"fmt"

	"github.com/samualhalder/lld/parkinglot_system/ticket"
	"github.com/samualhalder/lld/parkinglot_system/vehicle"
)

type Gates struct {
	Entry int
	Exit  int
}

type ParkingLot struct {
	// floors
	Floors map[int]*Floor
	Gates  Gates
	Rate   *Rate
	//
}

func (p *ParkingLot) AddFloor(floor *Floor) error {
	if _, ok := p.Floors[floor.Id]; ok {
		return fmt.Errorf("floor already exists")
	}
	p.Floors[floor.Id] = floor
	return nil
}

func (p *ParkingLot) ParkVehicle(vehicle vehicle.Vehicle) (*ticket.Ticket, error) {
	for _, floor := range p.Floors {
		tkt, err := floor.Park(vehicle)
		if err != nil {
			continue
		}
		return tkt, nil
	}
	return nil, fmt.Errorf("sorry cant park now")
}

func (p *ParkingLot) UnParkVehicle(vehicel vehicle.Vehicle) (*ticket.Ticket, error) {
	for _, floor := range p.Floors {
		tkt, ok := floor.UnPark(vehicel)
		if ok {
			vt := vehicel.IsType()
			r, err := p.Rate.GetRate(vt)
			if err != nil {
				return nil, err
			}
			dur := tkt.ExitTime.Sub(tkt.EntryTime).Seconds()
			amount := int(dur) * r
			tkt.Amount = amount
			return tkt, nil
		}
	}
	return nil, fmt.Errorf("no vehicle is found parked")
}
