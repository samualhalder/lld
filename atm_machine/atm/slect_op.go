package atm

import (
	"fmt"

	"github.com/samualhalder/lld/atm_machine/models"
)

type SelectOpr struct {
	BaseAtmSt
}



func (i *SelectOpr) InsertCard(crd *models.Card) error {
	// TODO: logic of Card inserted
	return nil
}
func (i *SelectOpr) Validate(pin int) error {
	return fmt.Errorf("NA")
}
func (i *SelectOpr) SelectOp(OpType) error {
	return nil
}
func (i *SelectOpr) WithDraw(amount int) error {
	return fmt.Errorf("NA")
}
func (i *SelectOpr) CheckBalance() (error, int) {
	return fmt.Errorf("NA"), 0
}
