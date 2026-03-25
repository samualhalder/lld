package loglevel

type LogLevel int

const (
	Info LogLevel = iota
	Debug
	Error
	Warn
	Panic
)

type LogLevelI interface {
	Log(LogLevel, string)
	Notify(string)
	Next() LogLevelI
}
