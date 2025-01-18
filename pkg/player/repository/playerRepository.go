package repository

import "github.com/kengkeng852/adventure-shop-api/entities"

type PlayerRepository interface{
	Creating(playerEntity *entities.Player) (*entities.Player, error)
	FindByID(playerID string) (*entities.Player, error)
}