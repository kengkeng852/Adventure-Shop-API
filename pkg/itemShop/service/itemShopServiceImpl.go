package service

import (
	"github.com/kengkeng852/adventure-shop-api/pkg/itemShop/repository"
	_itemShopModel "github.com/kengkeng852/adventure-shop-api/pkg/itemShop/model"   
)

type itemShopServiceImpl struct {
	itemShopRepository repository.ItemShopRepository
}

func NewItemShopServiceImpl(itemShopRepository repository.ItemShopRepository) ItemShopService {
	return &itemShopServiceImpl{itemShopRepository};
}

func (s *itemShopServiceImpl) Listing(itemFilter *_itemShopModel.ItemFilter) ([]*_itemShopModel.Item, error) {
	itemList, err := s.itemShopRepository.Listing(itemFilter)
	if err != nil {
		return nil, err
	}

	itemModelList := make([]*_itemShopModel.Item,0)
	for _, item := range itemList {
		itemModelList = append(itemModelList, item.ToItemModel())
	}
	
	return itemModelList, nil
} 