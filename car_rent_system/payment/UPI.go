package payment

type UpiPayment struct{}

func (c *UpiPayment) Pay(amout int) error {
	return nil
}
