package payment

import "fmt"

type Upi struct {
}

func (u *Upi) Pay(amount int) error {
	fmt.Printf("amount of %d is paid successfully using upi", amount)
	return nil
}
