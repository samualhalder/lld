package atm

import (
	"fmt"

	"github.com/samualhalder/lld/atm_machine/models"
)

type CheckBalance struct {
	// BaseAtmSt
	atm *Atm
}

func (i *CheckBalance) InsertCard(crd *models.Card) error {
	// TODO: logic of Card inserted
	return nil
}
func (i *CheckBalance) Validate(pin int) error {
	return fmt.Errorf("NA")
}
func (i *CheckBalance) SelectOp(OpType) error {
	return fmt.Errorf("NA")
}
func (i *CheckBalance) WithDraw(amount int) error {
	return fmt.Errorf("NA")
}
func (i *CheckBalance) CheckBalance() (int, error) {
	return i.atm.bankDB.GetBalance(i.atm.Card.CardNumber), nil

}
