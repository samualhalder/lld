package models

type Card struct {
	CardNumber int
	pin        int
}

func NewCard(cardNum, pin int) *Card {
	return &Card{
		CardNumber: cardNum,
		pin:        pin,
	}
}
