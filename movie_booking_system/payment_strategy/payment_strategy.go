package paymentstrategy

type PaymentStrategy interface {
	Pay(int)error
}
