package atm

import (
	"fmt"

	"github.com/samualhalder/lld/atm_machine/models"
)

type Idle struct {
	BaseAtmSt
}



func (i *Idle) InsertCard(crd *models.Card) error {
	// TODO: logic of Card inserted
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
func (i *Idle) CheckBalance() (error, int) {
	return fmt.Errorf("NA"), 0
}
