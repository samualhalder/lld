package payment

type Payment interface {
	Pay(int) error
}
