package repository

import (
	"user-service/internal/model"

	"gorm.io/gorm"
)

// UserRepository handles database operations related to users.
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new instance of UserRepository.
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

// GetAllUsers returns all users from the database.
func (ur *UserRepository) GetAllUsers() ([]model.User, error) {
	var users []model.User
	err := ur.db.Find(&users).Error
	return users, err
}

// GetUserByID returns a user by their ID from the database.
func (ur *UserRepository) GetUserByID(id int) (model.User, error) {
	var user model.User
	err := ur.db.First(&user, id).Error
	return user, err
}

// CreateUser creates a new user in the database.
func (ur *UserRepository) CreateUser(user model.User) (model.User, error) {
	err := ur.db.Create(&user).Error
	return user, err
}

// UpdateUser updates an existing user in the database.
func (ur *UserRepository) UpdateUser(updatedUser model.User) error {
	err := ur.db.Save(&updatedUser).Error
	return err
}

// DeleteUser deletes an existing user from the database by their ID.
func (ur *UserRepository) DeleteUser(id int) error {
	err := ur.db.Delete(&model.User{}, id).Error
	return err
}
