package models

type BalanceSheet struct {
	Transactions map[*User]float32
}
