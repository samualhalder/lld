package main

import (
	"github.com/samualhalder/lld/logger/loglevel"
)

type Logger struct {
	chainHead loglevel.LogLevelI
}

func NewLogger() *Logger {
	panic := loglevel.PanicLevel{}
	info := loglevel.InfoLevel{}
	debug := loglevel.DebugLevel{}
	warn := loglevel.WarnLevel{}
	error := loglevel.ErrorLevel{}
	info.NextLevel = &debug
	debug.NextLevel = &warn
	warn.NextLevel = &error
	error.NextLevel = &panic
	panic.NextLevel = nil
	logger := &Logger{
		chainHead: &info,
	}
	return logger
}

func (l *Logger) Info(msg string) {
	l.chainHead.Log(loglevel.Info, msg)
}
func (l *Logger) Debug(msg string) {
	l.chainHead.Log(loglevel.Debug, msg)
}

func (l *Logger) Warn(msg string) {
	l.chainHead.Log(loglevel.Warn, msg)
}

func (l *Logger) Error(msg string) {
	l.chainHead.Log(loglevel.Error, msg)
}

func (l *Logger) Panic(msg string) {
	l.chainHead.Log(loglevel.Panic, msg)
}
