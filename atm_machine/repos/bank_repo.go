package repos

import (
	"fmt"

	"github.com/samualhalder/lld/atm_machine/models"
)

type BankDB struct {
	accounts map[int]*models.Account

	cardToAccount map[int]int
}

func NewBankDB() *BankDB {
	return &BankDB{
		accounts:      make(map[int]*models.Account),
		cardToAccount: make(map[int]int),
	}
}

func (b *BankDB) GetAccountNumber(cardNumber int) int {
	return b.cardToAccount[cardNumber]
}

func (b *BankDB) GetBalance(accNum int) int {
	return b.accounts[accNum].CurrentBalance
}
func (b *BankDB) WithDraw(accNum, amount int) error {
	account, exist := b.accounts[accNum]
	if !exist {
		return fmt.Errorf("Wrong account number")
	}
	if account.CurrentBalance < amount {
		return fmt.Errorf("insufficient balance")
	}
	account.CurrentBalance -= amount
	return nil
}
