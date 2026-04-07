package denomination

import (
	"fmt"
	"maps"
)

type OneHundred struct {
	Next DenominationI
}

func (f *OneHundred) Handle(amount int, denominations map[DenominationType]int) (map[DenominationType]int, error) {
	usedDen := make(map[DenominationType]int)
	availableCnt := denominations[ONE_HUNDRED]

	req := amount / 100
	rem := amount % 100
	if availableCnt < req {
		canUse := req - availableCnt
		extra := req - canUse
		rem = rem + extra*100
		usedDen[ONE_HUNDRED] = canUse
	} else {
		usedDen[ONE_HUNDRED] = req
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
