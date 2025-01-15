package service

import (
	"github.com/kengkeng852/adventure-shop-api/entities"
	_itemShopModel "github.com/kengkeng852/adventure-shop-api/pkg/itemShop/model"
	"github.com/kengkeng852/adventure-shop-api/pkg/itemShop/repository"
)

type itemShopServiceImpl struct {
	itemShopRepository repository.ItemShopRepository
}

func NewItemShopServiceImpl(itemShopRepository repository.ItemShopRepository) ItemShopService {
	return &itemShopServiceImpl{itemShopRepository}
}

func (s *itemShopServiceImpl) Listing(itemFilter *_itemShopModel.ItemFilter) (*_itemShopModel.ItemResult, error) {
	itemList, err := s.itemShopRepository.Listing(itemFilter)
	if err != nil {
		return nil, err
	}

	itemCounting, err := s.itemShopRepository.Counting(itemFilter)
	if err != nil {
		return nil, err
	}

	size := itemFilter.Size
	page := itemFilter.Page
	totalPage := s.totalPageCalculation(itemCounting, size)

	return s.toItemResultResponse(itemList, page, totalPage), nil
}

func (s *itemShopServiceImpl) totalPageCalculation(totalItems int64, size int64) int64 {
	totalPage := totalItems / size

	if totalItems%size != 0 { //ex totalItems= 21 / size= 5  == 4 + 1(from totalPage++) == 5 pages
		totalPage++
	}

	return totalPage
}

func (s *itemShopServiceImpl) toItemResultResponse(itemEntityList []*entities.Item, page, totalPage int64) *_itemShopModel.ItemResult {
	itemModelList := make([]*_itemShopModel.Item, 0)

	for _, item := range itemEntityList {
		itemModelList = append(itemModelList, item.ToItemModel())
	}

	return &_itemShopModel.ItemResult{
		Items: itemModelList,
		Paginate: _itemShopModel.PaginateResult{
			Page:      page,
			TotalPage: totalPage,
		},
	}
}
