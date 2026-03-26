package loglevel

type InfoLevel struct {
	BaseLogLevel
}

func NewInfoLevel() *InfoLevel {
	return &InfoLevel{
		BaseLogLevel: *NewBaseLogLevel(),
	}
}

func (i *InfoLevel) Log(level LogLevel, msg string) {
	if level == Info {
		i.Notify(msg, "INFO")
		return
	}
	if i.Next() != nil {
		i.NextLevel.Log(level, msg)
	}
}
func (i *InfoLevel) Is(level LogLevel) bool {
	return level == Info
}
