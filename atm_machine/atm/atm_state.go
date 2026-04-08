package atm

import "github.com/samualhalder/lld/atm_machine/models"

type AtmState int

const (
	IdleSt AtmState = iota
	CardInsertedSt
	SelectOpSt
	WithdrawSt
	CheckBalanceSt
)

type OpType int

const (
	CheckBal OpType = iota
	WithdrawAmount
)

type AtmStateI interface {
	InsertCard(crd *models.Card) error
	Validate(pin int) error
	SelectOp(OpType) error
	WithDraw(amount int) error
	CheckBalance() (int, error)
}
