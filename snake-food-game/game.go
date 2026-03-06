package main

import "fmt"

type Directions int

const (
	Up Directions = iota
	Down
	Left
	Right
)

type Game struct {
	rows      int
	cols      int
	Snake     Snake
	Food      Food
	MaxScore  int
	CurrScore int
}

func (g *Game) Move(d Directions) error {
	switch d {
	case Up:
		g.Snake.MoveUp()
	case Down:
		g.Snake.MoveDown()
	case Left:
		g.Snake.MoveLeft()
	case Right:
		g.Snake.MoveRight()
	default:
		return fmt.Errorf("invalid move")
	}
	if err := g.ValideMove(); err != nil {
		g.GameOver()
		return err
	}
	return nil
}

func (g *Game) ValideMove() error {
	for _, b := range g.Snake.body {
		r := b.row
		c := b.col
		if r < 0 || r >= g.rows || c < 0 || c >= g.cols || g.Snake.HaveColided() {
			return fmt.Errorf("game over")
		}

	}
	head := g.Snake.body[0]
	if head == g.Food.pos {
		g.CurrScore++
		g.Food.GenerateFood(g.rows, g.cols)
	} else {
		g.Snake.ReductLength()
	}
	return nil
}

func (g *Game) GameOver() {
	if g.CurrScore > g.MaxScore {
		g.MaxScore = g.CurrScore
	}
	g.CurrScore = 0
}
