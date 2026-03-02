package board

import (
	"fmt"
	"strconv"
)

type Board struct {
	mat   [3][3]int
	ended bool
	who   int
}

func NewGame() *Board {
	var matrix [3][3]int

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			matrix[i][j] = -1
		}
	}
	return &Board{
		mat:   matrix,
		who:   0,
		ended: false,
	}
}

func (b *Board) Put(r, c int) (string, error) {
	println("who?", b.who)
	if b.ended {
		return "", fmt.Errorf("This game is over. Restart the game")
	}
	if r > 3 || c > 3 || r < 0 || c < 0 {
		return "", fmt.Errorf("give a valid board position")
	}
	if b.mat[r][c] != -1 {
		return "", fmt.Errorf("this is an used position")
	}
	b.mat[r][c] = b.who
	won := b.checkWin()
	if won {
		b.ended = true

		mssg := "player: " + strconv.Itoa(b.who) + " won the match"
		return mssg, nil
	}
	ended := b.isGameEned()
	if ended {
		b.ended = true
		return "match is ended", nil
	}
	if b.who == 1 {
		b.who = 0
	} else {
		b.who = 1
	}
	return "next player!", nil
}

func (b *Board) checkWin() bool {
	// check rows
	fl := true
	for i := 0; i <= 2; i++ {
		fl = true
		for j := 0; j <= 2; j++ {
			if b.mat[i][j] != b.who {
				fl = false
			}
		}
		if fl {

			return true
		}
	}

	// cols
	for i := 0; i <= 2; i++ {
		fl = true
		for j := 0; j <= 2; j++ {
			if b.mat[j][i] != b.who {
				fl = false
			}
		}
		if fl {
			if fl {

				return true
			}
			return true
		}
	}
	fl = true
	if b.mat[0][0] != b.who || b.mat[1][1] != b.who || b.mat[2][2] != b.who {
		fl = false
	}
	if fl {
		return true
	}
	fl = true
	if b.mat[0][2] != b.who || b.mat[1][1] != b.who || b.mat[2][0] != b.who {
		fl = false
	}
	if fl {
		return true
	}
	return false
}

func (b *Board) isGameEned() bool {
	for i := 0; i <= 2; i++ {
		for j := 0; j <= 2; j++ {
			if b.mat[i][j] == -1 {
				return false
			}
		}
	}
	return true
}

func (b *Board) Reset() {
	for i := 0; i <= 2; i++ {
		for j := 0; j <= 2; j++ {
			b.mat[i][j] = -1
		}
	}
	b.who = 0
	b.ended = false
}

func (b *Board) PrintBoard() {
	for i := range 3 {
		for j := range 3 {
			print(b.mat[i][j], " ")
		}
		print("\n")
	}
}
