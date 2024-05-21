package repository

import "user-service/internal/model"

// UserRepository defines the methods that any
// data storage provider needs to implement to get
// and store users.
type UserRepository interface {
	GetAllUsers() ([]model.User, error)
	GetUserByID(id int) (model.User, error)
	CreateUser(user model.User) (model.User, error)
	UpdateUser(user model.User) error
	DeleteUser(id int) error
}
