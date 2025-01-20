package repository

import (
	"github.com/kengkeng852/adventure-shop-api/entities"
	"gorm.io/gorm"
)

// _InventoryModel "github.com/kengkeng852/adventure-shop-api/pkg/inventory/model"

type InventoryRepository interface {
	Filling(tx *gorm.DB, playerID string, itemID uint64, qty int) ([]*entities.Inventory, error)
	Removing(tx *gorm.DB, playerID string, itemID uint64, qty int) error
	PlayerItemCounting(playerId string, itemID uint64) int64
	Listing(playerID string)([]*entities.Inventory, error)
}



