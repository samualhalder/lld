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
	quantity int
	price    int
	Category Category
}

func NewProduct(sku int, name string, price int) *Product {
	return &Product{
		SKU:      sku,
		Name:     name,
		price:    price,
		quantity: 0,
	}
}

func (p *Product) SetPrice(price int) {
	p.price = price
}

func (p *Product) AddQuantity(qn int) {
	p.quantity += qn
}

func (p *Product) RemQuantity(qn int) error {
	if qn > p.quantity {
		return fmt.Errorf("Stock not available to remove")
	}
	p.quantity -= qn
	return nil
}
