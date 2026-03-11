package models

type Direction int

const (
	Up Direction = iota
	Down
)

type Request struct {
	Floor *Floor
	Dir   Direction
}
