package appender

import (
	"fmt"
	"os"

	"github.com/samualhalder/lld/logger/formatter"
)

type FileAppender struct {
	Formatter formatter.Formatter
	filePath  string
}

func NewFileAppender(Formatter formatter.Formatter, filePath string) *FileAppender {
	return &FileAppender{
		Formatter: Formatter,
		filePath:  filePath,
	}
}

func (f *FileAppender) Append(text string, level string) error {
	forStr := f.Formatter.Format(text, level)
	str := fmt.Sprintf("%v", forStr)
	err := os.WriteFile(f.filePath, []byte(str), 0644)
	if err != nil {
		return err
	}
	return nil
}
