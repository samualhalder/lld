package loglevel

type WarnLevel struct {
	BaseLogLevel
}

func (i *WarnLevel) Log(level LogLevel, msg string) {
	if level == Warn {
		i.Notify(msg)
		return
	}
	if i.Next() != nil {
		i.NextLevel.Log(level, msg)
	}
}
