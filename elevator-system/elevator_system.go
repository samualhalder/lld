package main

import (
	"github.com/samualhalder/lld/elevator_system/models"
	"github.com/samualhalder/lld/elevator_system/schedule"
)

type ElevatorSystem struct {
	Floors   []*models.Floor
	Elevator []*models.Elevator
	schedule schedule.Schedule
}

func (e *ElevatorSystem) RequestElc(floor int, dir models.Direction) {
	req := &models.Request{
		Floor: e.Floors[floor],
		Dir:   dir,
	}
	elv := e.schedule.SelectElevator(req, e.Elevator)
	elv.AddRequest(req)
	if elv.State != models.Idle {
		elv.Move()
	}
}
