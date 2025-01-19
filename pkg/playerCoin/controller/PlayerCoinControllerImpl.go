package service

import (
	"net/http"

	"github.com/kengkeng852/adventure-shop-api/pkg/custom"
	_playerCoinModel "github.com/kengkeng852/adventure-shop-api/pkg/playerCoin/model"
	_playerCoinService "github.com/kengkeng852/adventure-shop-api/pkg/playerCoin/service"
	"github.com/kengkeng852/adventure-shop-api/pkg/validation"
	"github.com/labstack/echo/v4"
)

type playerCoinControllerImpl struct {
	playerCoinService _playerCoinService.PlayerCoinService
}

func NewPlayerCoinControllerImpl(playerCoinService _playerCoinService.PlayerCoinService) PlayerCoinController {
	return &playerCoinControllerImpl{playerCoinService}
}

func (c *playerCoinControllerImpl) CoinAdding(pctx echo.Context) error {
	playerID, err := validation.PlayerIDGetting(pctx)
	if err != nil {
		return custom.CustomError(pctx, http.StatusBadRequest, err.Error())
	}

	coinAddingReq := new(_playerCoinModel.CoinAddingReq)

	customEchoRequest := custom.NewCustomEchoRequest(pctx)

	if err := customEchoRequest.Bind(coinAddingReq); err != nil {
		return custom.CustomError(pctx, http.StatusBadRequest, err.Error())
	}
	coinAddingReq.PlayerID = playerID

	playerCoin, err := c.playerCoinService.CoinAdding(coinAddingReq)
	if err != nil {
		return custom.CustomError(pctx, http.StatusInternalServerError, err.Error())
	}

	return pctx.JSON(http.StatusCreated, playerCoin)
}

func (c *playerCoinControllerImpl) Showing(pctx echo.Context) error {
	playerID, err := validation.PlayerIDGetting(pctx)
	if err != nil {
		return custom.CustomError(pctx, http.StatusBadRequest, err.Error())
	}

	playerCoinShowing := c.playerCoinService.Showing(playerID)

	return pctx.JSON(http.StatusOK, playerCoinShowing)
}
