package models

type Card struct {
	CardNumber int
	Pin        int
}

func NewCard(cardNum, pin int) *Card {
	return &Card{
		CardNumber: cardNum,
		Pin:        pin,
	}
}
