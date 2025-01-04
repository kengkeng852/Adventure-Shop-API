package controller

import "github.com/kengkeng852/adventure-shop-api/pkg/itemShop/service"

type ItemShopControllerImpl struct {
	itemShopService service.ItemShopService
}

func NewItemShopControllerImpl(itemShopService service.ItemShopService) ItemShopController {
	return &ItemShopControllerImpl{itemShopService}
}
