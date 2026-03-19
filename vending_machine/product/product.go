package product

import "fmt"

type Product struct {
	Id       int
	Name     string
	Price    int
	Quantity int
}

func (p *Product) DecQuantity(num int) error {
	if p.Quantity < num {
		return fmt.Errorf("invalid quantity")
	}
	p.Quantity -= num
	return nil
}

