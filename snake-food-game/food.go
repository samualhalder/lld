package main

import "math/rand"

type Food struct {
	pos Cell
}

func (f *Food) GenerateFood(n, m int) {
	r := rand.Intn(n)
	c := rand.Intn(m)
	f.pos.col = c
	f.pos.row = r
}
