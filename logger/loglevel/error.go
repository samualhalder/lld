package loglevel

type ErrorLevel struct {
	BaseLogLevel
}

func NewErrorLevel() *ErrorLevel {
	return &ErrorLevel{
		BaseLogLevel: *NewBaseLogLevel(),
	}
}
func (i *ErrorLevel) Log(level LogLevel, msg string) {
	if level == Error {
		i.Notify(msg,"ERROR")
		return
	}
	if i.Next() != nil {
		i.NextLevel.Log(level, msg)
	}
}

func (i *ErrorLevel) Is(level LogLevel) bool {
	return level == Error
}
