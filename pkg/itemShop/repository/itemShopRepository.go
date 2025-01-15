package repository

import (
	"github.com/kengkeng852/adventure-shop-api/entities"
	_itemShopModel "github.com/kengkeng852/adventure-shop-api/pkg/itemShop/model"
)

type ItemShopRepository interface {
	Listing(itemFilter *_itemShopModel.ItemFilter) ([]*entities.Item, error)
}
