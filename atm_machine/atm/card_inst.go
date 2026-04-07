package atm

import (
	"fmt"

	"github.com/samualhalder/lld/atm_machine/models"
)

type CardInserted struct {
	BaseAtmSt
}

func (i *CardInserted) InsertCard(crd *models.Card) error {

	return fmt.Errorf("NA")
}
func (i *CardInserted) Validate(pin int) error {
	return nil
}
func (i *CardInserted) SelectOp(OpType) error {
	return fmt.Errorf("NA")
}
func (i *CardInserted) WithDraw(amount int) error {
	return fmt.Errorf("NA")
}
func (i *CardInserted) CheckBalance() (error, int) {
	return fmt.Errorf("NA"), 0
}
