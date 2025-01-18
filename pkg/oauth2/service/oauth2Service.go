package service

import (
	_playerModel "github.com/kengkeng852/adventure-shop-api/pkg/player/model"
	_adminModel "github.com/kengkeng852/adventure-shop-api/pkg/admin/model"
)

type OAuth2Service interface {
	PlayerAccountCreating(playerCreatingReq *_playerModel.PlayerCreatingReq) error
	AdminAccountCreating(adminCreatingReq *_adminModel.AdminCreatingReq) error
	IsPlayerValid(playerID string) bool 
	IsAdminValid(adminID string) bool 
}