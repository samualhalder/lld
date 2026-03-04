package player

import "github.com/samualhalder/lld/tic_tac_toe/board"

type Player struct {
	Name   string
	Symbol board.Symbol
}

func NewPlayer(name string, s board.Symbol) *Player {
	return &Player{
		Name:   name,
		Symbol: s,
	}
}
