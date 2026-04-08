package atm

import (
	"fmt"

	"github.com/samualhalder/lld/atm_machine/denomination"
	"github.com/samualhalder/lld/atm_machine/models"
	"github.com/samualhalder/lld/atm_machine/repos"
)

type Atm struct {
	id            int
	denominations map[denomination.DenominationType]int
	State         AtmStateI
	Card          *models.Card
	bankDB        *repos.BankDB
}

func NewAtm(id int, bankDb *repos.BankDB) *Atm {
	atm := &Atm{
		id:     id,
		bankDB: bankDb,
	}
	state := &Idle{
		atm: atm,
	}

	atm.State = state
	return atm
}

func (a *Atm) AddDenom(typ denomination.DenominationType, count int) {
	a.denominations[typ] += count
}

func (a *Atm) RemDenom(typ denomination.DenominationType, count int) error {
	currCnt := a.denominations[typ]
	if currCnt < count {
		return fmt.Errorf("insufficient denomination")
	}
	a.denominations[typ] -= count
	return nil
}

func (a *Atm) GetDenCnt(typ denomination.DenominationType) int {
	return a.denominations[typ]
}

func (a *Atm) InsertCard(card *models.Card) error {
	return a.State.InsertCard(card)
}

func (a *Atm) Verify(pin int) error {
	return a.State.Validate(pin)
}

func (a *Atm) SelectOperation(optype OpType) error {
	return a.State.SelectOp(optype)
}

func (a *Atm) CheckBal() (int, error) {
	return a.State.CheckBalance()
}

func (a *Atm) WithDraw(amount int) error {
	return a.State.WithDraw(amount)
}
