package service

import (
	"user-service/internal/model"
)

type UserService interface {
	CreateUser(name, email string) error
	GetUserByID(id uint) (*model.User, error)
}
