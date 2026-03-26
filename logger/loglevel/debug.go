package loglevel

type DebugLevel struct {
	BaseLogLevel
}

func NewDebugLevel() *DebugLevel {
	return &DebugLevel{
		BaseLogLevel: *NewBaseLogLevel(),
	}
}

func (i *DebugLevel) Log(level LogLevel, msg string) {
	if level == Debug {
		i.Notify(msg, "DEBUG")
		return
	}
	if i.Next() != nil {
		i.NextLevel.Log(level, msg)
	}
}

func (i *DebugLevel) Is(level LogLevel) bool {
	return level == Debug
}
