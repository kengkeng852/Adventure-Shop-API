package controller

import (
	"net/http"

	"github.com/kengkeng852/adventure-shop-api/pkg/custom"
	_inventoryService "github.com/kengkeng852/adventure-shop-api/pkg/inventory/service"
	"github.com/kengkeng852/adventure-shop-api/pkg/validation"
	"github.com/labstack/echo/v4"
)

type inventoryControllerImpl struct {
	inventoryService _inventoryService.InventoryService
	logger           echo.Logger
}

func NewinventoryControllerImpl(inventoryService _inventoryService.InventoryService, logger echo.Logger) InventoryController {
	return &inventoryControllerImpl{inventoryService, logger}
}

func (c *inventoryControllerImpl) Listing(pctx echo.Context) error {
	playerID, err := validation.PlayerIDGetting(pctx)
	if err != nil {
		c.logger.Errorf("error getting player id: %s", err.Error())
		return custom.CustomError(pctx, http.StatusBadRequest, err.Error())
	}

	inventoryListing, err := c.inventoryService.Listing(playerID)
	if err != nil {
		return custom.CustomError(pctx, http.StatusInternalServerError, err.Error())
	}

	return pctx.JSON(http.StatusOK, inventoryListing)
}
