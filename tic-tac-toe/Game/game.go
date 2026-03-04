package game

import (
	"fmt"

	"github.com/samualhalder/lld/tic_tac_toe/board"
	"github.com/samualhalder/lld/tic_tac_toe/player"
)

type Game struct {
	board     *board.Board
	players   [2]*player.Player
	scorecard map[*player.Player]int
	turn      *player.Player
	gameOver  bool
	ready     bool
}

func NewGame() *Game {
	return &Game{
		board:     board.NewBoard(),
		scorecard: make(map[*player.Player]int),
		gameOver:  false,
		ready:     false,
	}
}

func (g *Game) RegisterPlayer(player *player.Player) error {
	if g.players[0] == nil {
		g.players[0] = player
		g.scorecard[player] = 0
		return nil
	} else if g.players[1] == nil {
		sym := g.players[0].Symbol
		if sym == player.Symbol {
			return fmt.Errorf("This Symbol is already taken")
		}
		g.players[1] = player
		g.scorecard[player] = 0
		g.ready = true
		g.turn = g.players[0]
		return nil
	} else {
		return fmt.Errorf("two player already regester for this game!!!")
	}
}

func (g *Game) MakeMove(r, c int) error {
	if g.gameOver {
		return fmt.Errorf("This game is over, restart the game")
	}
	can := g.board.HaveEmptyCells()
	if !can {
		g.gameOver = true
		return fmt.Errorf("this game is over")
	}
	err := g.board.Put(r, c, g.turn.Symbol)
	if err != nil {
		return err
	}
	won := g.board.CheckWon(g.turn.Symbol)
	if won {
		g.scorecard[g.turn]++
		g.gameOver = true
		return nil
	}
	// TODO:  next players turn
	g.nextTurn()
	return nil
}

func (g *Game) nextTurn() {
	if g.turn == g.players[0] {
		g.turn = g.players[1]
	} else {
		g.turn = g.players[0]
	}
}

func (g *Game) RestartGame() {
	g.board.Reset()
	g.turn = g.players[0]
	g.gameOver = false
}

func (g *Game) GetScore() {
	println("Score for player ", g.players[0].Name, " is ", g.scorecard[g.players[0]])
	println("Score for player ", g.players[1].Name, " is ", g.scorecard[g.players[1]])
}
