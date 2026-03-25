package loglevel

type PanicLevel struct {
	BaseLogLevel
}

func (i *PanicLevel) Log(level LogLevel, msg string) {
	if level == Panic {
		i.Notify(msg)
		return
	}
	if i.Next() != nil {
		i.NextLevel.Log(level, msg)
	}
}
