package service

import (
	"github.com/kengkeng852/adventure-shop-api/entities"
	_adminModel "github.com/kengkeng852/adventure-shop-api/pkg/admin/model"
	_adminRepository "github.com/kengkeng852/adventure-shop-api/pkg/admin/repository"
	_playerModel "github.com/kengkeng852/adventure-shop-api/pkg/player/model"
	_playerRepository "github.com/kengkeng852/adventure-shop-api/pkg/player/repository"
)

type googleOAuth2Service struct {
	playerRepository _playerRepository.PlayerRepository
	adminRepository  _adminRepository.AdminRepository
}

func NewGoogleOAuth2Service(playerRepository _playerRepository.PlayerRepository, adminRepository _adminRepository.AdminRepository) *googleOAuth2Service {
	return &googleOAuth2Service{playerRepository, adminRepository}
}

func (s *googleOAuth2Service) PlayerAccountCreating(playerCreatingReq *_playerModel.PlayerCreatingReq) error {
	if !s.IsPlayerValid(playerCreatingReq.ID) { //shouldn't have at first so it can be created
		playerEntity := &entities.Player{
			ID:     playerCreatingReq.ID,
			Name:   playerCreatingReq.Name,
			Email:  playerCreatingReq.Email,
			Avatar: playerCreatingReq.Avatar,
		}

		if _, err := s.playerRepository.Creating(playerEntity); err != nil {
			return err
		}
	}

	return nil
}

func (s *googleOAuth2Service) AdminAccountCreating(adminCreatingReq *_adminModel.AdminCreatingReq) error {
	if !s.IsAdminValid(adminCreatingReq.ID) {
		adminEntity := &entities.Admin{
			ID:     adminCreatingReq.ID,
			Name:   adminCreatingReq.Name,
			Email:  adminCreatingReq.Email,
			Avatar: adminCreatingReq.Avatar,
		}

		if _, err := s.adminRepository.Creating(adminEntity); err != nil {
			return err
		}
	}

	return nil
}

func (s *googleOAuth2Service) IsPlayerValid(playerID string) bool {
	player, err := s.playerRepository.FindByID(playerID)
	if err != nil {
		return false
	}

	return player != nil
}

func (s *googleOAuth2Service) IsAdminValid(adminID string) bool {
	admin, err := s.adminRepository.FindByID(adminID)
	if err != nil {
		return false
	}

	return admin != nil
}
