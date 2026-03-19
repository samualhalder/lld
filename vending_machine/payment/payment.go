package payment

type PaymentStrategy interface {
	Pay(int) error
}
