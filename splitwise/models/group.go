package models

type Group struct {
	Id            int
	Name          string
	Participents  map[*User]*BalanceSheet
	Expences      []*Expense
}
