package service

import "github.com/kengkeng852/adventure-shop-api/pkg/itemShop/repository"

type itemShopServiceImpl struct {
	itemShopRepository repository.ItemShopRepository
}

func NewItemShopServiceImpl(itemShopRepository repository.ItemShopRepository) ItemShopService {
	return &itemShopServiceImpl{itemShopRepository};
}