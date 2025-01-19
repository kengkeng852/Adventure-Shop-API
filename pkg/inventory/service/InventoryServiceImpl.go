package service

import (
	"github.com/kengkeng852/adventure-shop-api/entities"
	_inventoryModel "github.com/kengkeng852/adventure-shop-api/pkg/inventory/model"
	_InventoryRepository "github.com/kengkeng852/adventure-shop-api/pkg/inventory/repository"
	_itemShopRepository "github.com/kengkeng852/adventure-shop-api/pkg/itemShop/repository"
)

type InventoryServiceImpl struct {
	InventoryRepository _InventoryRepository.InventoryRepository
	itemShopRepository  _itemShopRepository.ItemShopRepository
}

func NewInventoryServiceImpl(inventoryRepository _InventoryRepository.InventoryRepository, itemShopRepository _itemShopRepository.ItemShopRepository) InventoryService {
	return &InventoryServiceImpl{inventoryRepository, itemShopRepository}
}

func (s *InventoryServiceImpl) Listing(playerID string) ([]*_inventoryModel.Inventory, error) {
	inventoryEntities, err := s.InventoryRepository.Listing(playerID)
	if err != nil {
		return nil, err
	}

	uniqueItemWithQuantitiyCounterList := s.getUniqueItemWithQuantityCounterList(inventoryEntities)

	return s.buildInventoryListingResult(uniqueItemWithQuantitiyCounterList), nil
}

func (s *InventoryServiceImpl) getUniqueItemWithQuantityCounterList(inventoryEntities []*entities.Inventory) []_inventoryModel.ItemQuantityCounting {
	itemQuantityCounterList := make([]_inventoryModel.ItemQuantityCounting, 0)

	itemMapWithQuantity := make(map[uint64]uint)

	for _, inventory := range inventoryEntities {
		itemMapWithQuantity[inventory.ItemID]++ //itemMapWithQuantity[1] = 1, after ++ = itemMapWithQuantity[1] = 2
	}

	for itemID, quantity := range itemMapWithQuantity {
		itemQuantityCounterList = append(itemQuantityCounterList, _inventoryModel.ItemQuantityCounting{
			ItemID:   itemID,
			Quantity: quantity,
		})
	}

	return itemQuantityCounterList
}

func (s *InventoryServiceImpl) buildInventoryListingResult(
	uniqueItemWithQuantitiyCounterList []_inventoryModel.ItemQuantityCounting,
) []*_inventoryModel.Inventory {
	uniqueItemIDList := s.getItemID(uniqueItemWithQuantitiyCounterList)

	itemEntities, err := s.itemShopRepository.FindByIDList(uniqueItemIDList)
	if err != nil {
		return make([]*_inventoryModel.Inventory, 0)
	}

	results := make([]*_inventoryModel.Inventory, 0)
	itemMapWithQuantity := s.getItemMapWithQuantity(uniqueItemWithQuantitiyCounterList)

	for _, itemEntity := range itemEntities {
		results = append(results, &_inventoryModel.Inventory{
			Item:     itemEntity.ToItemModel(),
			Quantity: itemMapWithQuantity[itemEntity.ID],
		})
	}
	return results
}

func (s *InventoryServiceImpl) getItemID(
	uniqueItemWithQuantitiyCounterList []_inventoryModel.ItemQuantityCounting,
) []uint64 {
	uniqueItemIDList := make([]uint64, 0)

	for _, inventory := range uniqueItemWithQuantitiyCounterList {
		uniqueItemIDList = append(uniqueItemIDList, inventory.ItemID)
	}

	return uniqueItemIDList
}

func (s *InventoryServiceImpl) getItemMapWithQuantity(
	uniqueItemWithQuantitiyCounterList []_inventoryModel.ItemQuantityCounting,
) map[uint64]uint {
	itemMapWithQuantity := make(map[uint64]uint)

	for _, inventory := range uniqueItemWithQuantitiyCounterList {
		itemMapWithQuantity[inventory.ItemID] = inventory.Quantity
	}

	return itemMapWithQuantity
}
