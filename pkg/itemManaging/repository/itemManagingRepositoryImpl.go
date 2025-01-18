package repository

import (
	"github.com/kengkeng852/adventure-shop-api/databases"
	"github.com/kengkeng852/adventure-shop-api/entities"
	_itemManagingException "github.com/kengkeng852/adventure-shop-api/pkg/itemManaging/exception"
	_itemManagingModel "github.com/kengkeng852/adventure-shop-api/pkg/itemManaging/model"
	"github.com/labstack/echo/v4"
)

type itemManagingRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewItemManagingRepositoryImpl(db databases.Database, logger echo.Logger) *itemManagingRepositoryImpl {
	return &itemManagingRepositoryImpl{db, logger}
}

func (r *itemManagingRepositoryImpl) Creating(itemEntity *entities.Item) (*entities.Item, error) {
	item := new(entities.Item) //create to take data from itemEntity

	if err := r.db.Connect().Create(itemEntity).Scan(item).Error; err != nil {
		r.logger.Errorf("Creating item failed: %s", err.Error())
		return nil, &_itemManagingException.ItemCreating{}
	}

	return item, nil
}

func (r *itemManagingRepositoryImpl) Editing(itemID uint64, itemEditingReq *_itemManagingModel.ItemEditingReq) (uint64, error) {
	if err := r.db.Connect().Model(&entities.Item{}).Where("id = ?", itemID).Updates(itemEditingReq).Error; err != nil {
		r.logger.Errorf("Editing item failed: %s", err.Error())
		return 0, &_itemManagingException.ItemEditing{}
	}

	return itemID, nil
}

func (r *itemManagingRepositoryImpl) Archiving(itemID uint64) error {
	if err := r.db.Connect().Table("items").Where("id = ?", itemID).Update("is_archive", true).Error; err != nil {
		r.logger.Errorf("Archiving item failed: %s", err.Error())
		return &_itemManagingException.ItemArchiving{ItemID: itemID}
	}

	return nil
}
