package service

import (
	_inventoryModel "github.com/kengkeng852/adventure-shop-api/pkg/inventory/model"
)

type InventoryService interface {
	Listing(playerID string) ([]*_inventoryModel.Inventory, error)
}
