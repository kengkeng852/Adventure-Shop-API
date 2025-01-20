package repository

import (
	"github.com/kengkeng852/adventure-shop-api/databases"
	"github.com/kengkeng852/adventure-shop-api/entities"
	_inventoryException "github.com/kengkeng852/adventure-shop-api/pkg/inventory/exception"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type inventoryRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewInventoryRepositoryImpl(db databases.Database, logger echo.Logger) InventoryRepository {
	return &inventoryRepositoryImpl{db, logger}
}

func (r *inventoryRepositoryImpl) Filling(tx *gorm.DB, playerID string, itemID uint64, qty int) ([]*entities.Inventory, error) {
	connection := r.db.Connect()
	if tx != nil {
		connection = tx
	}

	inventoryEntities := make([]*entities.Inventory, 0)

	for range qty {
		inventoryEntities = append(inventoryEntities, &entities.Inventory{
			PlayerID: playerID,
			ItemID: itemID,
		})
	}

	if err := connection.Create(inventoryEntities).Error; err != nil {
		r.logger.Errorf("error filling inventory: %s", err)
		return nil, &_inventoryException.InventoryFilling{
			PlayerID: playerID,
			ItemID:   itemID,
		}
	}

	return inventoryEntities, nil
}

func (r *inventoryRepositoryImpl) Removing(tx *gorm.DB, playerID string, itemID uint64, qty int) error {
	connection := r.db.Connect()
	if tx != nil {
		connection = tx
	}

	inventoryEntites, err := r.findPlayerItemInInventoryByID(playerID, itemID, qty)
	if err != nil {
		return err
	}

	for _, inventory := range inventoryEntites {
		inventory.IsDeleted = true

		if err := connection.Model(&entities.Inventory{}).Where("id = ?", inventory.ID).Updates(inventory).Error; err != nil {
			tx.Rollback()
			r.logger.Errorf("error removing player item in inventory: %s", err)
			return &_inventoryException.PlayerItemRemoving{ItemID: itemID}
		}
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

func (r *inventoryRepositoryImpl) findPlayerItemInInventoryByID(playerID string, itemID uint64, qty int) ([]*entities.Inventory, error) {
	inventoryEntities := make([]*entities.Inventory, 0)

	if err := r.db.Connect().Where(
		"player_id = ? and item_id = ? and is_deleted = ?", playerID, itemID, false,
	).Limit(qty).Find(&inventoryEntities).Error; err != nil {
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
