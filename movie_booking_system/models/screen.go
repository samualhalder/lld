package models

type Screen struct {
	Id    string
	Seats map[string]*Seat
	Movie *Movie
}

func NewScreen(id string) *Screen {
	return &Screen{
		Id:    id,
		Seats: make(map[string]*Seat),
	}
}

func (s *Screen) AddSeat(seat *Seat) {
	s.Seats[seat.Id] = seat
}
