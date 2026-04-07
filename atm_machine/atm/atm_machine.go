package atm

import (
	"fmt"

	"github.com/samualhalder/lld/atm_machine/denomination"
	"github.com/samualhalder/lld/atm_machine/models"
)

type Atm struct {
	id            int
	denominations map[denomination.DenominationType]int
	State         AtmStateI
	Card          *models.Card
}

func NewAtm(id int) *Atm {
	return &Atm{
		id:    id,
		State: &Idle{},
	}
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
