package order

import "github.com/samualhalder/lld/vending_mechine_system/product"

type OrderState int

const (
	Pending OrderState = iota
	Failed
	PaymentFailure
	Completed
)

type Order struct {
	Id          int
	Products    map[*product.Product]int
	TotalAmount int
	Discount    int
	State       OrderState
}

func (o Order) AddProduct(p *product.Product) {
	_, ok := o.Products[p]
	if !ok {
		o.Products[p] = 1
	} else {
		o.Products[p]++
	}
}
func (o Order) RemProduct(p *product.Product) {
	val := o.Products[p]
	if val == 0 {
		delete(o.Products, p)
	} else {
		o.Products[p]--
	}
}
