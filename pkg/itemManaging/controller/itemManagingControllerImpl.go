package controller

import (
	"net/http"
	"strconv"

	"github.com/kengkeng852/adventure-shop-api/pkg/custom"
	_itemManagingModel "github.com/kengkeng852/adventure-shop-api/pkg/itemManaging/model"
	_itemManagingService "github.com/kengkeng852/adventure-shop-api/pkg/itemManaging/service"
	"github.com/kengkeng852/adventure-shop-api/pkg/validation"
	"github.com/labstack/echo/v4"
)

type itemManagingControllerImpl struct {
	itemManagingService _itemManagingService.ItemManagingService
}

func NewItemManagingControllerImpl(itemManagingService _itemManagingService.ItemManagingService) ItemManagingController {
	return &itemManagingControllerImpl{itemManagingService}
}

func (c *itemManagingControllerImpl) Creating(pctx echo.Context) error {
	adminID, err := validation.AdminIDGetting(pctx)
	if err != nil {
		return custom.CustomError(pctx, http.StatusBadRequest, err.Error())
	}

	itemCreatingReq := new(_itemManagingModel.ItemCreatingReq)

	customEchoRequest := custom.NewCustomEchoRequest(pctx)

	if err := customEchoRequest.Bind(itemCreatingReq); err != nil {
		return custom.CustomError(pctx, http.StatusBadRequest, err.Error())
	}
	itemCreatingReq.AdminID = adminID

	item, err := c.itemManagingService.Creating(itemCreatingReq)
	if err != nil {
		return custom.CustomError(pctx, http.StatusInternalServerError, err.Error())
	}

	return pctx.JSON(http.StatusCreated, item)
}

func (c *itemManagingControllerImpl) Editing(pctx echo.Context) error {
	itemID, err := getItemID(pctx)
	if err != nil {
		return custom.CustomError(pctx, http.StatusBadRequest, err.Error())
	}

	itemEditingReq := new(_itemManagingModel.ItemEditingReq)

	customEchoRequest := custom.NewCustomEchoRequest(pctx)

	if err := customEchoRequest.Bind(itemEditingReq); err != nil {
		return custom.CustomError(pctx, http.StatusBadRequest, err.Error())
	}

	item, err := c.itemManagingService.Editing(itemID, itemEditingReq)
	if err != nil {
		return custom.CustomError(pctx, http.StatusInternalServerError, err.Error())
	}

	return pctx.JSON(http.StatusOK, item)
}

func getItemID(pctx echo.Context) (uint64, error) {
	itemID := pctx.Param("itemID")
	itemIDUint64, err := strconv.ParseUint(itemID, 10, 64)
	if err != nil {
		return 0, nil
	}
	return itemIDUint64, nil
}

func (c *itemManagingControllerImpl) Archiving(pctx echo.Context) error {
	itemID, err := getItemID(pctx)
	if err != nil {
		return custom.CustomError(pctx, http.StatusBadRequest, err.Error())
	}

	if err := c.itemManagingService.Archiving(itemID); err != nil {
		return custom.CustomError(pctx, http.StatusInternalServerError, err.Error())
	}

	return pctx.NoContent(http.StatusNoContent)
}
