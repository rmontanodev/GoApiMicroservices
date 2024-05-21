package repository

import (
	"user-service/internal/model"

	"gorm.io/gorm"
)

// UserRepositoryImpl handles database operations related to users.
type UserRepositoryImpl struct {
	db *gorm.DB
}

// NewUserRepositoryImpl creates a new instance of UserRepositoryImpl.
func NewUserRepositoryImpl(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db}
}

// GetAllUsers returns all users from the database.
func (ur *UserRepositoryImpl) GetAllUsers() ([]model.User, error) {
	var users []model.User
	err := ur.db.Find(&users).Error
	return users, err
}

// GetUserByID returns a user by their ID from the database.
func (ur *UserRepositoryImpl) GetUserByID(id int) (model.User, error) {
	var user model.User
	err := ur.db.First(&user, id).Error
	return user, err
}

// CreateUser creates a new user in the database.
func (ur *UserRepositoryImpl) CreateUser(user model.User) (model.User, error) {
	err := ur.db.Create(&user).Error
	return user, err
}

// UpdateUser updates an existing user in the database.
func (ur *UserRepositoryImpl) UpdateUser(updatedUser model.User) error {
	err := ur.db.Save(&updatedUser).Error
	return err
}

// DeleteUser deletes an existing user from the database by their ID.
func (ur *UserRepositoryImpl) DeleteUser(id int) error {
	err := ur.db.Delete(&model.User{}, id).Error
	return err
}
