package denomination

type DenominationType int

const (
	FIVE_HUNDRED DenominationType = iota
	TWO_HUNDRED
	ONE_HUNDRED
)

type DenominationI interface {
	Handle(amount int, denominations map[DenominationType]int) error
	MoveNext()
}
