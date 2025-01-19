package repository

import "github.com/kengkeng852/adventure-shop-api/entities"

// _InventoryModel "github.com/kengkeng852/adventure-shop-api/pkg/inventory/model"

type InventoryRepository interface {
	Filling(inventoryEntities []*entities.Inventory) ([]*entities.Inventory, error)
	Removing(playerID string, itemID uint64, limit int) error
	PlayerItemCounting(playerId string, itemID uint64) int64
	Listing(playerID string)([]*entities.Inventory, error)
}
