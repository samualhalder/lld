package models

import "time"

type Operation int

const (
	Add Operation = iota
	Remove
)

type Log struct {
	ProductId     int
	OperationType Operation
	Timestamp     time.Time
}

func NewLog(id int, ty Operation) *Log {
	return &Log{
		ProductId:     id,
		OperationType: ty,
		Timestamp:     time.Now(),
	}
}
