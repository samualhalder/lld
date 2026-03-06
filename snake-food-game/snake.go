package main

type Snake struct {
	body []Cell
}

func (s *Snake) MoveUp() {
	n := len(s.body)
	head := s.body[n-1]
	s.body = append(s.body, Cell{head.row - 1, head.col})
}

func (s *Snake) MoveDown() {
	n := len(s.body)
	head := s.body[n-1]
	s.body = append(s.body, Cell{head.row + 1, head.col})
}

func (s *Snake) MoveLeft() {
	n := len(s.body)
	head := s.body[n-1]
	s.body = append(s.body, Cell{head.row, head.col - 1})
}

func (s *Snake) MoveRight() {
	n := len(s.body)
	head := s.body[n-1]
	s.body = append(s.body, Cell{head.row, head.col + 1})
}
func (s *Snake) ReductLength() {
	s.body = s.body[1:]
}

func (s *Snake) HaveColided() bool {
	mp := make(map[Cell]int)
	for _, b := range s.body {
		_, ok := mp[b]
		if ok {
			return true
		}
		mp[b] = 1
	}
	return false
}
