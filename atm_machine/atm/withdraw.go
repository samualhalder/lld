package atm

import (
	"fmt"

	"github.com/samualhalder/lld/atm_machine/denomination"
	"github.com/samualhalder/lld/atm_machine/models"
)

type WithDraw struct {
	// BaseAtmSt
	atm *Atm
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
	accNum := i.atm.bankDB.GetAccountNumber(i.atm.Card.CardNumber)
	bal:=i.atm.bankDB.GetBalance(accNum)
	if bal<amount{
		return fmt.Errorf("insufficient balance")
	}
	_,err:= i.atm.SelectDenominations.Handle(amount,i.atm.denominations)
	if err!=nil{
		return err
	}
	
	return nil
}
func (i *WithDraw) CheckBalance() (int, error) {
	return 0, fmt.Errorf("NA")
}
