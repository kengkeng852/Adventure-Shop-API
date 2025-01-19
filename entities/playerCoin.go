package entities

import ("time" 
_playerCoinModel "github.com/kengkeng852/adventure-shop-api/pkg/playerCoin/model"
)
  
type PlayerCoin struct {  
	ID        uint64 `gorm:"primaryKey;autoIncrement;"`  
	PlayerID  string `gorm:"type:varchar(64);not null;"`  
	Amount    int64  `gorm:"not null;"`  
	CreatedAt time.Time `gorm:"not null;autoCreateTime;"`  
}

func (c *PlayerCoin) ToPlayerCoinModel() *_playerCoinModel.PlayerCoin{
	return &_playerCoinModel.PlayerCoin{
		ID: c.ID,   
		PlayerID: c.PlayerID,
		Amount: c.Amount,
		CreatedAt: c.CreatedAt,
	}
}