package denomination

import (
	"fmt"
	"maps"
)

type TwoHundred struct {
	Next DenominationI
}

func (f *TwoHundred) Handle(amount int, denominations map[DenominationType]int) (map[DenominationType]int, error) {
	usedDen := make(map[DenominationType]int)
	availableCnt := denominations[TWO_HUNDRED]

	req := amount / 200
	rem := amount % 200
	if availableCnt < req {
		canUse := req - availableCnt
		extra := req - canUse
		rem = rem + extra*200
		usedDen[TWO_HUNDRED] = canUse
	} else {
		usedDen[TWO_HUNDRED] = req
	}

	if f.Next == nil && rem > 0 {
		return nil, fmt.Errorf("insufficient notes available in machine")
	}

	if rem == 0 {
		return usedDen, nil
	}

	otherDen, err := f.Next.Handle(rem, denominations)
	if err != nil {
		return nil, err
	}
	maps.Copy(usedDen, otherDen)
	return usedDen, nil
}
