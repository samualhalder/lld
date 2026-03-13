package alert

type Alert interface {
	MakeAllert(quantity int) error
}
