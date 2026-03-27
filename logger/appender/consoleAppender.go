package appender

import (
	"fmt"

	"github.com/samualhalder/lld/logger/formatter"
)

type ConsoleAppender struct {
	Formatter formatter.Formatter
}

func NewConsoleAppender(formatter formatter.Formatter) *ConsoleAppender {
	return &ConsoleAppender{
		Formatter: formatter,
	}
}

func (c *ConsoleAppender) Append(messg string, level string) error {
	// TODO: first format the string then
	fmt.Println(c.Formatter.Format(messg, level))
	return nil
}
