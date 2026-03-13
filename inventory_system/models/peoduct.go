package models

import "fmt"

type Category int

const (
	Electronics Category = iota
	Food
	Dress
)

type Product struct {
	SKU      int
	Name     string
	Quantity int
	price    int
	Category Category
}

func NewProduct(sku int, name string, price int) *Product {
	return &Product{
		SKU:      sku,
		Name:     name,
		price:    price,
		Quantity: 0,
	}
}

func (p *Product) SetPrice(price int) {
	p.price = price
}

func (p *Product) AddQuantity(qn int) {
	p.Quantity += qn
}

func (p *Product) RemQuantity(qn int) error {
	if qn > p.Quantity {
		return fmt.Errorf("Stock not available to remove")
	}
	p.Quantity -= qn
	return nil
}
