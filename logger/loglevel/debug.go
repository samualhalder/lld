package loglevel

type DebugLevel struct {
	BaseLogLevel
}

func (i *DebugLevel) Log(level LogLevel, msg string) {
	if level == Debug {
		i.Notify(msg)
		return
	}
	if i.Next() != nil {
		i.NextLevel.Log(level, msg)
	}
}
