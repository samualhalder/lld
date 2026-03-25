package loglevel

type InfoLevel struct {
	BaseLogLevel
}

func (i *InfoLevel) Log(level LogLevel, msg string) {
	if level == Info {
		i.Notify(msg)
		return
	}
	if i.Next() != nil {
		i.NextLevel.Log(level, msg)
	}
}
