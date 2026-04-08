package atm

import (
	"fmt"

	"github.com/samualhalder/lld/atm_machine/models"
)

type Idle struct {
	// BaseAtmSt
	atm *Atm
}

func (i *Idle) InsertCard(crd *models.Card) error {
	i.atm.Card = crd
	i.atm.State = &Validate{atm: i.atm}
	return nil
}
func (i *Idle) Validate(pin int) error {
	return fmt.Errorf("NA")
}
func (i *Idle) SelectOp(OpType) error {
	return fmt.Errorf("NA")
}
func (i *Idle) WithDraw(amount int) error {
	return fmt.Errorf("NA")
}
func (i *Idle) CheckBalance() (int,error) {
	return 0,fmt.Errorf("NA")
}
