package warehouse

import (
	"fmt"

	"github.com/samualhalder/lld/inventory_system/models"
)

type WareHouse struct {
	ID       int
	products map[int]*models.Product
	// alert
}

func NewWareHouse(id int) *WareHouse {
	return &WareHouse{
		ID:       id,
		products: make(map[int]*models.Product),
	}
}

func (w *WareHouse) Add(id, queantity int) (error, *models.Log) {
	pr, ok := w.products[id]
	if !ok {
		return fmt.Errorf("No product is stroed in stock"), nil
	}
	pr.AddQuantity(queantity)
	log := models.NewLog(pr.SKU, models.Add)
	return nil, log

}

func (w *WareHouse) Rem(id, queantity int) (error, *models.Log) {
	pr, ok := w.products[id]
	if !ok {
		return fmt.Errorf("No product is stroed in stock"), nil
	}
	err := pr.RemQuantity(queantity)
	if err != nil {
		return err, nil
	}
	log := models.NewLog(pr.SKU, models.Remove)
	return nil, log

}

func (w *WareHouse) GetProduct(sku int) *models.Product {
	pr, ok := w.products[sku]
	if !ok {
		return nil
	}
	return pr
}

func (w *WareHouse) List() []*models.Product {
	var products []*models.Product
	for _, pr := range w.products {
		products = append(products, pr)
	}
	return products
}

func (w *WareHouse) Entry(product *models.Product) error {
	_, ok := w.products[product.SKU]
	if ok {
		return fmt.Errorf("this prosuct already exists in warehouse")

	}
	w.products[product.SKU] = product
	return nil
}
