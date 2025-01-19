package repository

import (
	"github.com/kengkeng852/adventure-shop-api/databases"
	"github.com/kengkeng852/adventure-shop-api/entities"
	_inventoryException "github.com/kengkeng852/adventure-shop-api/pkg/inventory/exception"
	"github.com/labstack/echo/v4"
)

type inventoryRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewInventoryRepositoryImpl(db databases.Database, logger echo.Logger) InventoryRepository {
	return &inventoryRepositoryImpl{db, logger}
}

func (r *inventoryRepositoryImpl) Filling(inventoryEntities []*entities.Inventory) ([]*entities.Inventory, error) {
	inventoryEntitiesResult := make([]*entities.Inventory, 0)

	if err := r.db.Connect().CreateInBatches(inventoryEntities, len(inventoryEntities)).Scan(&inventoryEntitiesResult).Error; err != nil {
		r.logger.Errorf("error filling inventory: %s", err)
		return nil, &_inventoryException.InventoryFilling{
			PlayerID: inventoryEntities[0].PlayerID,
			ItemID:   inventoryEntities[0].ItemID,
		}
	}

	return inventoryEntitiesResult, nil
}

func (r *inventoryRepositoryImpl) Removing(playerID string, itemID uint64, limit int) error {
	inventoryEntites, err := r.findPlayerItemInInventoryByID(playerID, itemID, limit)
	if err != nil {
		return err
	}

	tx := r.db.Connect().Begin() //start transaction

	for _, inventory := range inventoryEntites {
		inventory.IsDeleted = true

		if err := tx.Model(&entities.Inventory{}).Where("id = ?", inventory.ID).Updates(inventory).Error; err != nil {
			tx.Rollback()
			r.logger.Errorf("error removing player item in inventory: %s", err)
			return &_inventoryException.PlayerItemRemoving{ItemID: itemID}
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback() //if err while commit so rollback
		r.logger.Errorf("error removing player item in inventory: %s", err)
		return &_inventoryException.PlayerItemRemoving{ItemID: itemID}
	}

	return nil
}

func (r *inventoryRepositoryImpl) Listing(playerID string) ([]*entities.Inventory, error) {
	inventoryEntities := make([]*entities.Inventory, 0)

	if err := r.db.Connect().Where(
		"player_id = ? and is_deleted = ?", playerID, false,
	).Find(&inventoryEntities).Error; err != nil {
		r.logger.Errorf("error listing player inventory: %s", err)
		return nil, &_inventoryException.PlayerItemListing{PlayerID: playerID}
	}

	return inventoryEntities, nil
}

func (r *inventoryRepositoryImpl) findPlayerItemInInventoryByID(playerID string, itemID uint64, limit int) ([]*entities.Inventory, error) {
	inventoryEntities := make([]*entities.Inventory, 0)

	if err := r.db.Connect().Where(
		"player_id = ? and item_id = ? and is_deleted = ?", playerID, itemID, false,
	).Limit(limit).Find(&inventoryEntities).Error; err != nil {
		r.logger.Errorf("error finding player item in inventory by ID: %s", err)
		return nil, &_inventoryException.PlayerItemFinding{ItemID: itemID}
	}

	return inventoryEntities, nil
}

func (r *inventoryRepositoryImpl) PlayerItemCounting(playerID string, itemID uint64) int64 {
	var itemCount int64

	if err := r.db.Connect().Model(&entities.Inventory{}).Where("player_id = ? and item_id = ? and is_deleted = ?", playerID, itemID, false).Count(&itemCount).Error; err != nil {
		r.logger.Errorf("error counting player item in inventory: %s", err)
		return -1
	}

	return itemCount
}
