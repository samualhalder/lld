package formatter

import (
	"time"
)

type TextFormatter struct{}

func NewTextFormatter() *TextFormatter {
	return &TextFormatter{}
}

func (t *TextFormatter) Format(msg string, level string) any {
	time := time.Now().UTC().String()

	return time + " [" + level + "] " + msg
}
