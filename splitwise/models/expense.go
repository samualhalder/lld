package models

import "github.com/samualhalder/lld/splitwise/enums"

type Expense struct {
	Desc          string
	Amount        float32
	PaidBy        *User
	Parties       map[*User]float32
	SplitStrategy enums.SplitType
}
