package loglevel

type ErrorLevel struct {
	BaseLogLevel
}

func (i *ErrorLevel) Log(level LogLevel, msg string) {
	if level == Debug {
		i.Notify(msg)
		return
	}
	if i.Next() != nil {
		i.NextLevel.Log(level, msg)
	}
}
