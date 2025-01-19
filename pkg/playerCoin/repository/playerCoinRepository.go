package repository

import (
	"github.com/kengkeng852/adventure-shop-api/entities"
	_playerCoinModel "github.com/kengkeng852/adventure-shop-api/pkg/playerCoin/model"
)

type PlayerCoinRepository interface {
	CoinAdding(playerCoinEntity *entities.PlayerCoin) (*entities.PlayerCoin, error)
	Showing(playerID string) (*_playerCoinModel.PlayerCoinShowing,error)
}
