package repository

import "github.com/kengkeng852/adventure-shop-api/entities"

type AdminRepository interface{
	Creating(adminEntity *entities.Admin) (*entities.Admin, error)
	FindByID(adminID string) (*entities.Admin, error)
}