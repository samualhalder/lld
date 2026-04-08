package atm

import (
	"fmt"

	"github.com/samualhalder/lld/atm_machine/models"
)

type Validate struct {
	BaseAtmSt
	atm *Atm
}

func (i *Validate) InsertCard(crd *models.Card) error {

	return fmt.Errorf("NA")
}
func (i *Validate) Validate(pin int) error {
	if i.atm.Card.Pin != pin {
		return fmt.Errorf("wrogn pin")
	}
	i.atm.State = &SelectOpr{atm: i.atm}
	return nil
}
func (i *Validate) SelectOp(OpType) error {
	return fmt.Errorf("NA")
}
func (i *Validate) WithDraw(amount int) error {
	return fmt.Errorf("NA")
}
func (i *Validate) CheckBalance() (int, error) {
	return 0, fmt.Errorf("NA")
}
