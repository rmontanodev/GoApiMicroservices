package service

import (
	"user-service/internal/model"
	"user-service/internal/repository"
)

type UserServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserServiceImpl(userRepository repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{userRepository: userRepository}
}

func (s *UserServiceImpl) CreateUser(name, email string) (model.User, error) {
	user := &model.User{
		Name:  name,
		Email: email,
	}
	return s.userRepository.CreateUser(*user)
}

func (s *UserServiceImpl) GetUserByID(id int) (model.User, error) {
	return s.userRepository.GetUserByID(id)
}
