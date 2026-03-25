package loglevel

type BaseLogLevel struct {
	Appenders []int
	NextLevel LogLevelI
}

func (b *BaseLogLevel) Notify(message string) {
	for _, app := range b.Appenders {
		app += 1 // TODO: have to call append method of each appender
	}
}

func (b *BaseLogLevel) Next() LogLevelI {
	return b.NextLevel
}
