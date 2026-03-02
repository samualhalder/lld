package main

import (
	"github.com/samualhalder/lld/tic_tac_toe/board"
)

func main() {
	game1 := board.NewGame()
	game1.Put(0, 0)
	game1.Put(0, 1)
	game1.Put(1, 0)
	game1.Put(0, 2)
	msf, _ := game1.Put(2, 0)
	print(msf)
	game1.PrintBoard()
	game1.Reset()
	game1.PrintBoard()
}
