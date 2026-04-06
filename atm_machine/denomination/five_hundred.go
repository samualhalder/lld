package denomination

type FiveHundred struct {
	Next DenominationI
}

func (f *FiveHundred) Handle(amount int, denominations map[DenominationType]int) error {
	avinvleCount := denominations[FIVE_HUNDRED]
	req := amount / 500
	rem := amount % 500
	if rem == 0 {
		return nil
	}
	return f.Next.Handle(rem, denominations)
}
