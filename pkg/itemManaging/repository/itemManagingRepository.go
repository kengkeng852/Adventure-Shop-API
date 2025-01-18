package repository

import (
	"github.com/kengkeng852/adventure-shop-api/entities"
	_itemManagingModel "github.com/kengkeng852/adventure-shop-api/pkg/itemManaging/model"
)

type ItemManagingRepository interface {
	Creating(itemEntity *entities.Item) (*entities.Item, error)
	Editing(itemID uint64, itemEditingReq *_itemManagingModel.ItemEditingReq) (uint64, error)
	Archiving(itemID uint64) error
}
