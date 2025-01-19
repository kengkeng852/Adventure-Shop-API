package service

import (
	"github.com/kengkeng852/adventure-shop-api/entities"
	_playerCoinModel "github.com/kengkeng852/adventure-shop-api/pkg/playerCoin/model"
	_playerCoinRepository "github.com/kengkeng852/adventure-shop-api/pkg/playerCoin/repository"
)

type playerCoinServiceImpl struct {
	playerCoinRepository _playerCoinRepository.PlayerCoinRepository
}

func NewPlayerCoinServiceImpl(playerCoinRepository _playerCoinRepository.PlayerCoinRepository) PlayerCoinService {
	return &playerCoinServiceImpl{playerCoinRepository}
}

func (s *playerCoinServiceImpl) CoinAdding(coinAddingReq *_playerCoinModel.CoinAddingReq) (*_playerCoinModel.PlayerCoin, error) {
	playerCoinEntity := entities.PlayerCoin{
		PlayerID: coinAddingReq.PlayerID,
		Amount:   coinAddingReq.Amount,
	}

	playerCoinEntityResult, err := s.playerCoinRepository.CoinAdding(&playerCoinEntity)
	if err != nil {
		return nil, err
	}

	playerCoinEntityResult.PlayerID = coinAddingReq.PlayerID

	return playerCoinEntityResult.ToPlayerCoinModel(), nil
}

func (s *playerCoinServiceImpl) Showing(playerID string) (*_playerCoinModel.PlayerCoinShowing) {
	playerCoinShowing, err := s.playerCoinRepository.Showing(playerID)
	if err != nil {
		return &_playerCoinModel.PlayerCoinShowing{
			PlayerID: playerID,
			Coin: 0,
		}
	}

	return playerCoinShowing
}
