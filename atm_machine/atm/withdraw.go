package atm

import (
	"fmt"

	"github.com/samualhalder/lld/atm_machine/models"
)

type WithDraw struct {
	BaseAtmSt
}

func (i *WithDraw) InsertCard(crd *models.Card) error {
	// TODO: logic of Card inserted
	return nil
}
func (i *WithDraw) Validate(pin int) error {
	return fmt.Errorf("NA")
}
func (i *WithDraw) SelectOp(OpType) error {
	return fmt.Errorf("NA")
}
func (i *WithDraw) WithDraw(amount int) error {
	return nil
}
func (i *WithDraw) CheckBalance() (error, int) {
	return fmt.Errorf("NA"), 0
}
