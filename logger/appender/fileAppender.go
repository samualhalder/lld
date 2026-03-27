package appender

import (
	"fmt"
	"os"
	"sync"

	"github.com/samualhalder/lld/logger/formatter"
)

type FileAppender struct {
	Formatter formatter.Formatter
	filePath  string
	mu        *sync.RWMutex
}

func NewFileAppender(Formatter formatter.Formatter, filePath string) *FileAppender {
	return &FileAppender{
		Formatter: Formatter,
		filePath:  filePath,
	}
}

func (f *FileAppender) Append(text string, level string) error {
	f.mu.RLocker()
	formatt := f.Formatter.Format(text, level)
	str := fmt.Sprintf("%v\n", formatt)
	file, err := os.OpenFile(f.filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	file.Write([]byte(str))
	f.mu.RLocker().Unlock()
	return nil
}
