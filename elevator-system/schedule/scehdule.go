package schedule

import "github.com/samualhalder/lld/elevator_system/models"

type Schedule interface {
	SelectElevator(req *models.Request, elvs []*models.Elevator) *models.Elevator
}
