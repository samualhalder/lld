package schedule

import (
	"github.com/samualhalder/lld/elevator_system/models"
)

type ClosestElv struct{}

func (c *ClosestElv) SelectElevator(req *models.Request, elvs []*models.Elevator) *models.Elevator {
	var elevator *models.Elevator

	for _, el := range elvs {
		if elevator == nil {
			elevator = el
		} else {
			if el.CurrentFloor < elevator.CurrentFloor {
				elevator = el
			}
		}
	}
	return elevator
}
