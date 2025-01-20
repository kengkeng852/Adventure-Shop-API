package controller

import (
	"net/http"

	"github.com/kengkeng852/adventure-shop-api/pkg/custom"
	_itemShopModel "github.com/kengkeng852/adventure-shop-api/pkg/itemShop/model"
	_itemShopService "github.com/kengkeng852/adventure-shop-api/pkg/itemShop/service"
	"github.com/kengkeng852/adventure-shop-api/pkg/validation"
	"github.com/labstack/echo/v4"
)

type itemShopControllerImpl struct {
	itemShopService _itemShopService.ItemShopService
}

func NewItemShopControllerImpl(itemShopService _itemShopService.ItemShopService) ItemShopController {
	return &itemShopControllerImpl{itemShopService}
}

func (c *itemShopControllerImpl) Listing(pctx echo.Context) error {
	itemFilter := new(_itemShopModel.ItemFilter)

	customEchoRequest := custom.NewCustomEchoRequest(pctx)

	if err := customEchoRequest.Bind(itemFilter); err != nil {
		return custom.CustomError(pctx, http.StatusBadRequest, err.Error())
	}

	itemModelList, err := c.itemShopService.Listing(itemFilter)
	if err != nil {
		return custom.CustomError(pctx, http.StatusInternalServerError, err.Error())
	}

	return pctx.JSON(http.StatusOK, itemModelList)
}

func (c *itemShopControllerImpl) Buying(pctx echo.Context) error {
	playerID, err := validation.PlayerIDGetting(pctx)
	if err != nil {
		return custom.CustomError(pctx, http.StatusBadRequest, err.Error())
	}

	buyingReq:= new(_itemShopModel.BuyingReq)

	customEchoRequest := custom.NewCustomEchoRequest(pctx)

	if err := customEchoRequest.Bind(buyingReq); err != nil {
		return custom.CustomError(pctx, http.StatusBadRequest, err.Error())
	}

	buyingReq.PlayerID = playerID

	playerCoin, err := c.itemShopService.Buying(buyingReq)
	if err != nil {
		return custom.CustomError(pctx, http.StatusInternalServerError, err.Error())
	}

	return pctx.JSON(http.StatusOK, playerCoin)
}

func (c *itemShopControllerImpl) Selling(pctx echo.Context) error {
	playerID, err := validation.PlayerIDGetting(pctx)
	if err != nil {
		return custom.CustomError(pctx, http.StatusBadRequest, err.Error())
	}

	sellingReq:= new(_itemShopModel.SellingReq)

	customEchoRequest := custom.NewCustomEchoRequest(pctx)

	if err := customEchoRequest.Bind(sellingReq); err != nil {
		return custom.CustomError(pctx, http.StatusBadRequest, err.Error())
	}

	sellingReq.PlayerID = playerID

	playerCoin, err := c.itemShopService.Selling(sellingReq)
	if err != nil {
		return custom.CustomError(pctx, http.StatusInternalServerError, err.Error())
	}

	return pctx.JSON(http.StatusOK, playerCoin)
}
