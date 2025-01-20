package repository

import (
	"github.com/kengkeng852/adventure-shop-api/entities"
	_itemShopModel "github.com/kengkeng852/adventure-shop-api/pkg/itemShop/model"
	"gorm.io/gorm"
)

type ItemShopRepository interface {
	TransactionBegin() *gorm.DB
	TransactionRollback(tx *gorm.DB) error
	TransactionCommit(tx *gorm.DB) error
	Listing(itemFilter *_itemShopModel.ItemFilter) ([]*entities.Item, error)
	Counting(itemFilter *_itemShopModel.ItemFilter) (int64, error)
	FindByID(itemID uint64) (*entities.Item, error)
	FindByIDList(itemIDList []uint64) ([]*entities.Item, error)
	PurchaseHistoryRecording(tx *gorm.DB,purchasingEntity *entities.PurchaseHistory) (*entities.PurchaseHistory, error)
}
