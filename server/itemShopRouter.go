package server

import (
	_inventoryRepository "github.com/kengkeng852/adventure-shop-api/pkg/inventory/repository"
	_itemShopController "github.com/kengkeng852/adventure-shop-api/pkg/itemShop/controller"
	_itemShopRepository "github.com/kengkeng852/adventure-shop-api/pkg/itemShop/repository"
	_itemShopService "github.com/kengkeng852/adventure-shop-api/pkg/itemShop/service"
	_playerCoinRepository "github.com/kengkeng852/adventure-shop-api/pkg/playerCoin/repository"
)

func (s *echoServer) initItemShopRouter(m *authorizingMiddleware) {
	router := s.app.Group("/v1/item-shop")

	itemShopRepository := _itemShopRepository.NewItemShopRepositoryImpl(s.db, s.app.Logger)
	playerCoinRepository := _playerCoinRepository.NewPlayerCoinRepositoryImpl(s.db, s.app.Logger)
	inventoryRepository := _inventoryRepository.NewInventoryRepositoryImpl(s.db, s.app.Logger)

	itemShopService := _itemShopService.NewItemShopServiceImpl(
		itemShopRepository,
		playerCoinRepository,
		inventoryRepository,
		s.app.Logger)

	itemShopController := _itemShopController.NewItemShopControllerImpl(itemShopService)

	router.GET("", itemShopController.Listing)
	router.POST("/buying", itemShopController.Buying, m.PlayerAuthorizing)
	router.POST("/selling", itemShopController.Selling, m.PlayerAuthorizing)
}
