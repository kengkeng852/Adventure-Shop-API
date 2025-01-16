package repository

import (
	"github.com/kengkeng852/adventure-shop-api/entities"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	_itemShopException "github.com/kengkeng852/adventure-shop-api/pkg/itemShop/exception"
	_itemShopModel "github.com/kengkeng852/adventure-shop-api/pkg/itemShop/model" 
)

type itemShopRepositoryImpl struct {
	db     *gorm.DB
	logger echo.Logger
}

func NewItemShopRepositoryImpl(db *gorm.DB, logger echo.Logger) ItemShopRepository {
	return &itemShopRepositoryImpl{db, logger}
}

func (r *itemShopRepositoryImpl) Listing(itemFilter *_itemShopModel.ItemFilter) ([]*entities.Item, error) {
	itemList := make([]*entities.Item, 0)

	//filter
	query := r.db.Model(&entities.Item{}).Where("is_archive = ?", false)// select * from items

	if itemFilter.Name != "" {
		query = query.Where("name ilike ?", "%"+itemFilter.Name+"%") // use ?: by variable
	}
	if itemFilter.Description != "" {
		query = query.Where("description ilike ?", "%"+itemFilter.Description+"%")
	}

	//pagination
	// item 1 2 3 4 5 6 7 8 9 10
	//      0       | 5 index
	// if want item at page 2
	// (2 - 1) * size:5 = 5 so it start at index 5 and show limit item per page
	offset := int((itemFilter.Page-1) * itemFilter.Size) //start with which product
	limit := int(itemFilter.Size)

	if err := query.Offset(offset).Limit(limit).Find(&itemList).Error; err != nil {
		r.logger.Errorf("Failed to list items: %s", err.Error())
		return nil, &_itemShopException.ItemListing{} //use to call specific error 
	}

	return itemList, nil
}


func (r *itemShopRepositoryImpl) Counting(itemFilter *_itemShopModel.ItemFilter) (int64, error) {
	//filter
	query := r.db.Model(&entities.Item{}).Where("is_archive = ?", false)// select * from items

	if itemFilter.Name != "" {
		query = query.Where("name ilike ?", "%"+itemFilter.Name+"%") // use ?: by variable
	}
	if itemFilter.Description != "" {
		query = query.Where("description ilike ?", "%"+itemFilter.Description+"%")
	}

	var count int64

	if err := query.Count(&count).Error; err != nil {
		r.logger.Errorf("Counting item failed: %s", err.Error())
		return -1, &_itemShopException.ItemCounting{} //use to call specific error 
	}

	return count, nil
}

func (r *itemShopRepositoryImpl) FindByID(itemID uint64) (*entities.Item, error) {
	item := new(entities.Item)

	if err := r.db.First(item, itemID).Error; err != nil { //First() is find that specific on the first data that found
		r.logger.Errorf("Failed to find item by ID: %s", err.Error())
		return nil, &_itemShopException.ItemNotFound{}
	}

	return item,nil
}
