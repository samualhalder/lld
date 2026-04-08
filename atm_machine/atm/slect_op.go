package atm

import (
	"fmt"

	"github.com/samualhalder/lld/atm_machine/models"
)

type SelectOpr struct {
	BaseAtmSt
	atm *Atm
}

func (i *SelectOpr) InsertCard(crd *models.Card) error {
	// TODO: logic of Card inserted
	return nil
}
func (i *SelectOpr) Validate(pin int) error {
	return fmt.Errorf("NA")
}
func (i *SelectOpr) SelectOp(opr OpType) error {
	if opr == CheckBal {
		i.atm.State = &CheckBalance{atm: i.atm}
	} else if opr == WithdrawAmount {
		i.atm.State = &WithDraw{atm: i.atm}
	} else {
		return fmt.Errorf("Unknown operation")
	}
	return nil
}
func (i *SelectOpr) WithDraw(amount int) error {
	return fmt.Errorf("NA")
}
func (i *SelectOpr) CheckBalance() (int, error) {
	return 0, fmt.Errorf("NA")
}
