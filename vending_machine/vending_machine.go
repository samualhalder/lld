package main

import (
	"fmt"

	"github.com/samualhalder/lld/vending_mechine_system/order"
	"github.com/samualhalder/lld/vending_mechine_system/payment"
	"github.com/samualhalder/lld/vending_mechine_system/product"
)

type MachineState int

const (
	Idle MachineState = iota
	Processing
	Maintanance
	PaymentPending
)

type VendingMachine struct {
	stock        map[*product.Product]int
	state        MachineState
	currentOrder *order.Order
	orders       []*order.Order
	payment      payment.PaymentStrategy
}

func (v *VendingMachine) GerProduct() map[*product.Product]int {
	return v.stock
}

func (v *VendingMachine) AddToCart(p *product.Product) error {
	v.state = Processing
	var currrentOrder *order.Order
	if v.currentOrder == nil {
		currrentOrder = &order.Order{}
	} else {
		currrentOrder = v.currentOrder
	}
	qnt := v.stock[p]
	if qnt == 0 {
		return fmt.Errorf("out of stock")
	}
	currrentOrder.AddProduct(p)
	return nil
}

func (v *VendingMachine) RemoveFromCart(p *product.Product) error {
	var currrentOrder *order.Order
	if v.currentOrder == nil {
		return fmt.Errorf("no product to remove")
	} else {
		currrentOrder = v.currentOrder
	}
	qnt, ok := v.stock[p]
	if qnt == 0 || !ok {
		return fmt.Errorf("no product to remove")
	}
	currrentOrder.RemProduct(p)
	return nil
}

func (v *VendingMachine) cancelOrder() {
	v.state = Idle
	v.currentOrder = nil
}

func (v *VendingMachine) PlaceOrder() *order.Order {
	currOrder := v.currentOrder
	v.currentOrder = nil
	v.state = PaymentPending
	return currOrder
}
func (v *VendingMachine) CompleteOrder(order *order.Order) {
	for pro, qty := range order.Products {
		v.stock[pro] -= qty
	}
}

func (v *VendingMachine) PayAndProcess(paySt payment.PaymentStrategy, order *order.Order) error {
	err := paySt.Pay(order.TotalAmount)
	if err != nil {
		return fmt.Errorf("payment processing")
	}
	v.CompleteOrder(order)
	return nil
}
