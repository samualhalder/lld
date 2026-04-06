package main

import (
	"github.com/samualhalder/lld/logger/appender"
	"github.com/samualhalder/lld/logger/formatter"
	"github.com/samualhalder/lld/logger/loglevel"
)

func main() {
	logger := NewLogger()
	logger.Subscribe(loglevel.Info, appender.NewConsoleAppender(formatter.NewTextFormatter()))
	logger.Subscribe(loglevel.Debug, appender.NewConsoleAppender(formatter.NewTextFormatter()))
	logger.Subscribe(loglevel.Error, appender.NewConsoleAppender(formatter.NewTextFormatter()))
	logger.Subscribe(loglevel.Error, appender.NewFileAppender(formatter.NewTextFormatter(), "logs.txt"))
	logger.Subscribe(loglevel.Info, appender.NewFileAppender(formatter.NewTextFormatter(), "logs.txt"))

	logger.Debug("its a debug")
	logger.Info("Its a test comment")
	logger.Warn("its a warning")
	logger.Error("Null pointer exception")
}
