package models

type Account struct {
	AccountNumber     int
	AccountHolderName string
	AtmCardNo         int
	CurrentBalance    int
}

func NewAccount(name string, number int) *Account {
	return &Account{
		AccountNumber:     number,
		AccountHolderName: name,
	}
}
