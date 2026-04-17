package models

import "time"

type Movie struct {
	id      int
	Name    string
	Runtime time.Duration
}
