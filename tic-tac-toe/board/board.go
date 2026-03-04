package board

import "fmt"

type Symbol int

const (
	X Symbol = iota
	O
	Empty
)

type Board struct {
	mat [3][3]Symbol
}

func NewBoard() *Board {
	var matrix [3][3]Symbol

	for i := range 3 {
		for j := range 3 {
			matrix[i][j] = Empty
		}
	}
	return &Board{
		mat: matrix,
	}
}

func (b *Board) Put(r, c int, s Symbol) error {
	if r < 0 || r >= 3 || c < 0 || c >= 3 {
		return fmt.Errorf("illigal move: out of bound")
	}
	if b.mat[r][c] != Empty {
		return fmt.Errorf("illigal move: already filled")
	}
	b.mat[r][c] = s
	return nil
}

func (b *Board) CheckWon(s Symbol) bool {
	// row wise
	for i := range 3 {
		fl := true
		for j := range 3 {
			if b.mat[i][j] != s {
				fl = false
				break
			}
		}
		if fl {
			return true
		}
	}

	for i := range 3 {
		fl := true
		for j := range 3 {
			if b.mat[j][i] != s {
				fl = false
				break
			}
		}
		if fl {
			return true
		}
	}
	fl := true
	for i := 0; i < 3; i++ {
		if b.mat[i][i] != s {
			fl = false
			break
		}
	}
	if fl {
		return true
	}
	fl = true
	for i := 0; i < 3; i++ {
		for j := 2; j >= 0; j-- {
			if b.mat[i][j] != s {
				fl = false
				break
			}
		}
	}
	return fl
}

func (b *Board) Reset() {
	for i := range len(b.mat) {
		for j := range len(b.mat[i]) {
			b.mat[i][j] = Empty
		}
	}
}

func (b *Board) HaveEmptyCells() bool {
	return true
}

// func (b *Board) checkWin() bool {
// 	// check rows
// 	fl := true
// 	for i := 0; i <= 2; i++ {
// 		fl = true
// 		for j := 0; j <= 2; j++ {
// 			if b.mat[i][j] != b.who {
// 				fl = false
// 			}
// 		}
// 		if fl {

// 			return true
// 		}
// 	}

// 	// cols
// 	for i := 0; i <= 2; i++ {
// 		fl = true
// 		for j := 0; j <= 2; j++ {
// 			if b.mat[j][i] != b.who {
// 				fl = false
// 			}
// 		}
// 		if fl {
// 			if fl {

// 				return true
// 			}
// 			return true
// 		}
// 	}
// 	fl = true
// 	if b.mat[0][0] != b.who || b.mat[1][1] != b.who || b.mat[2][2] != b.who {
// 		fl = false
// 	}
// 	if fl {
// 		return true
// 	}
// 	fl = true
// 	if b.mat[0][2] != b.who || b.mat[1][1] != b.who || b.mat[2][0] != b.who {
// 		fl = false
// 	}
// 	if fl {
// 		return true
// 	}
// 	return false
// }

// func (b *Board) isGameEned() bool {
// 	for i := 0; i <= 2; i++ {
// 		for j := 0; j <= 2; j++ {
// 			if b.mat[i][j] == -1 {
// 				return false
// 			}
// 		}
// 	}
// 	return true
// }

// func (b *Board) Reset() {
// 	for i := 0; i <= 2; i++ {
// 		for j := 0; j <= 2; j++ {
// 			b.mat[i][j] = -1
// 		}
// 	}
// 	b.who = 0
// 	b.ended = false
// }

// func (b *Board) PrintBoard() {
// 	for i := range 3 {
// 		for j := range 3 {
// 			print(b.mat[i][j], " ")
// 		}
// 		print("\n")
// 	}
// }
