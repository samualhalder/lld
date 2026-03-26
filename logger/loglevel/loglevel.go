package loglevel

import "github.com/samualhalder/lld/logger/appender"

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
	Notify(string, string)
	Next() LogLevelI
	Is(LogLevel) bool
	Subs(appender.Appernder)
}
