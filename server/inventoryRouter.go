package server

import (
	_inventoryController "github.com/kengkeng852/adventure-shop-api/pkg/inventory/controller"
	_inventoryRepository "github.com/kengkeng852/adventure-shop-api/pkg/inventory/repository"
	_inventoryService "github.com/kengkeng852/adventure-shop-api/pkg/inventory/service"
	_itemShopRepository "github.com/kengkeng852/adventure-shop-api/pkg/itemShop/repository"
)

func (s *echoServer) initInventoryRouter(m *authorizingMiddleware) {
	router := s.app.Group("/v1/inventory")

	inventoryRepository := _inventoryRepository.NewInventoryRepositoryImpl(s.db, s.app.Logger)
	itemShopRepository := _itemShopRepository.NewItemShopRepositoryImpl(s.db,s.app.Logger)

	inventoryService := _inventoryService.NewInventoryServiceImpl(inventoryRepository,itemShopRepository)

	inventoryController := _inventoryController.NewinventoryControllerImpl(inventoryService, s.app.Logger)

	router.GET("", inventoryController.Listing, m.PlayerAuthorizing)
}