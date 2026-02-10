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
	//
}

func (p *ParkingLot) AddFloor(id int, floor *Floor) error {
	if _, ok := p.Floors[id]; ok {
		return fmt.Errorf("floor already exists")
	}
	p.Floors[id] = floor
	return nil
}

func (p *ParkingLot) ParkVehicle(vehicle vehicle.Vehicle) (*ticket.Ticket, error) {
	for _, floor := range p.Floors {
		tkt, err := floor.Park(vehicle)
		if err != nil {
			return nil, err
		}
		return tkt, nil
	}
	return nil, fmt.Errorf("sorry cant park now")
}
