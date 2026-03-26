package loglevel

type PanicLevel struct {
	BaseLogLevel
}
func NewPanicLevel() *PanicLevel {
	return &PanicLevel{
		BaseLogLevel: *NewBaseLogLevel(),
	}
}
func (i *PanicLevel) Log(level LogLevel, msg string) {
	if level == Panic {
		i.Notify(msg,"PANIC")
		return
	}
	if i.Next() != nil {
		i.NextLevel.Log(level, msg)
	}
}

func (i *PanicLevel) Is(level LogLevel) bool {
	return level == Panic
}
