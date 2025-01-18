package service

import (
	_itemManagingModel "github.com/kengkeng852/adventure-shop-api/pkg/itemManaging/model"
	_itemShopModel "github.com/kengkeng852/adventure-shop-api/pkg/itemShop/model"
)

type ItemManagingService interface {
	Creating(itemCreatingReq *_itemManagingModel.ItemCreatingReq) (*_itemShopModel.Item, error) // return itemShopModel cause this model intended to call for showing item
	Editing(itemID uint64, itemEditingReq *_itemManagingModel.ItemEditingReq) (*_itemShopModel.Item, error)
	Archiving(itemID uint64) error
}
