package storage

type Storage interface {
	Get(string) *string
	Put(string, string)
}
