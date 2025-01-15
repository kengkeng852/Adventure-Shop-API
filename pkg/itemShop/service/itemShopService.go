package service

import _itemShopModel "github.com/kengkeng852/adventure-shop-api/pkg/itemShop/model"  


type ItemShopService interface {
	Listing(itemFilter *_itemShopModel.ItemFilter) (*_itemShopModel.ItemResult, error)
}