package storage

import "time"

type Storage interface {
	Get(string) *string
	Put(string, string, time.Duration)
}
