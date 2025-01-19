package model

import _itemShopModel "github.com/kengkeng852/adventure-shop-api/pkg/itemShop/model"

type (
	Inventory struct {
		Item     *_itemShopModel.Item `json:"item"`
		Quantity uint                 `json:"quantity"`
	}

	ItemQuantityCounting struct {
		ItemID   uint64
		Quantity uint
	}
)
