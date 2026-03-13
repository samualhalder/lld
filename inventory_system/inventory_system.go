package main

import (
	"fmt"

	"github.com/samualhalder/lld/inventory_system/models"
	"github.com/samualhalder/lld/inventory_system/warehouse"
)

type InventorySystem struct {
	warehouses []*warehouse.WareHouse
	histroy    *models.History
}

func (i *InventorySystem) AddProduct(warehouseId, sku, quantity int) error {
	if warehouseId > len(i.warehouses) {
		return fmt.Errorf("ho such warehouse")
	}
	warehouse := i.warehouses[warehouseId]
	err, log := warehouse.Add(sku, quantity)
	if err != nil {
		return err
	}
	i.histroy.Add(log)
	return nil
}

func (i *InventorySystem) RemProduct(warehouseId, sku, quantity int) error {
	if warehouseId > len(i.warehouses) {
		return fmt.Errorf("ho such warehouse")
	}
	warehouse := i.warehouses[warehouseId]
	err, log := warehouse.Rem(sku, quantity)
	if err != nil {
		return err
	}
	i.histroy.Add(log)
	return nil
}

func (i *InventorySystem) GetHistory() []*models.Log {
	return i.histroy.List()
}
