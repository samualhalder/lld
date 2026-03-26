package loglevel

type WarnLevel struct {
	BaseLogLevel
}

func NewWarnLevel() *WarnLevel {
	return &WarnLevel{
		BaseLogLevel: *NewBaseLogLevel(),
	}
}

func (i *WarnLevel) Log(level LogLevel, msg string) {
	if level == Warn {
		i.Notify(msg, "WARN")
		return
	}
	if i.Next() != nil {
		i.NextLevel.Log(level, msg)
	}
}

func (i *WarnLevel) Is(level LogLevel) bool {
	return level == Warn
}
