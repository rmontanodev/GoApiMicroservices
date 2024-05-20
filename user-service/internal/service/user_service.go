package service

import (
	"user-service/internal/model"
	"user-service/internal/repository"
)

type UserService interface {
	CreateUser(name, email string) error
	GetUserByID(id uint) (*model.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository: userRepository}
}

func (s *userService) CreateUser(name, email string) error {
	user := &model.User{
		Name:  name,
		Email: email,
	}
	return s.userRepository.Create(user)
}

func (s *userService) GetUserByID(id uint) (*model.User, error) {
	return s.userRepository.GetByID(id)
}
