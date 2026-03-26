package loglevel

import "github.com/samualhalder/lld/logger/appender"

type BaseLogLevel struct {
	Appenders map[appender.Appernder]int
	NextLevel LogLevelI
}

func NewBaseLogLevel() *BaseLogLevel {
	return &BaseLogLevel{
		Appenders: make(map[appender.Appernder]int),
		NextLevel: nil,
	}
}

func (b *BaseLogLevel) Notify(message string, level string) {
	for app, _ := range b.Appenders {
		app.Append(message, level)
	}
}

func (b *BaseLogLevel) Next() LogLevelI {
	return b.NextLevel
}

func (b *BaseLogLevel) Subs(appender appender.Appernder) {
	b.Appenders[appender] = 1
}
