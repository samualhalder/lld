package models

type State int

const (
	Idle State = iota
	GoingUp
	GoingDown
)

type Elevator struct {
	Id              int
	State           State
	CurrentFloor    int
	Direction       Direction
	MaxFloors       int
	internalRequest []int
	externalRequest []*Request
}

func (e *Elevator) AddRequest(Req *Request) {
	e.externalRequest = append(e.externalRequest, Req)
}
func (e *Elevator) AddStop(fl int) {
	e.internalRequest = append(e.internalRequest, fl)
}
func (e *Elevator) Move() {
	// for len(e.externalRequest) > 0 || len(e.internalRequest) > 0 {

	// }
}
