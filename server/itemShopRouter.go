package server

import (
	_itemShopController "github.com/kengkeng852/adventure-shop-api/pkg/itemShop/controller"
	_itemShopRepository "github.com/kengkeng852/adventure-shop-api/pkg/itemShop/repository"
	_itemShopService "github.com/kengkeng852/adventure-shop-api/pkg/itemShop/service"
)

func (s *echoServer) initItemShopRouter() {
	router := s.app.Group("/v1/item-shop")

	itemShopRepository := _itemShopRepository.NewItemShopRepositoryImpl(s.db, s.app.Logger)
	itemShopService := _itemShopService.NewItemShopServiceImpl(itemShopRepository)
	itemShopController := _itemShopController.NewItemShopControllerImpl(itemShopService)

	router.GET("", itemShopController.Listing)
}
