package payment

import "fmt"

type CoinPayment struct {
}

func (u *CoinPayment) Pay(amount int) error {
	fmt.Printf("amount of %d is paid successfully in coins", amount)
	return nil
}
