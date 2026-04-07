package denomination

type DenominationType int

const (
	FIVE_HUNDRED DenominationType = iota
	TWO_HUNDRED
	ONE_HUNDRED
)

type DenominationI interface {
	Handle(amount int, denominations map[DenominationType]int) (map[DenominationType]int, error)
}

func NewDenomicataionController() DenominationI {
	five := &FiveHundred{}
	two := &TwoHundred{}
	one := &OneHundred{}
	five.Next = two
	two.Next = one
	return five
}
