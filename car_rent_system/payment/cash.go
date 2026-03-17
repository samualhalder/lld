package payment

type CashPayment struct{}

func (c *CashPayment) Pay(amout int) error {
	return nil
}
