package repository_test

import (
	"log"
	"testing"
	"user-service/internal/config"
	"user-service/internal/model"
	"user-service/internal/repository"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	cfg := config.LoadConfig()

	dsn := cfg.DatabaseURL

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Clear the users table before each test
	db.Exec("TRUNCATE TABLE users RESTART IDENTITY CASCADE")

	return db
}

func TestCreateUserRepo(t *testing.T) {
	db := setupTestDB()
	repo := repository.NewUserRepositoryImpl(db)

	user := model.User{Name: "Test User", Email: "test@example.com"}
	createdUser, err := repo.CreateUser(user)

	assert.NoError(t, err)
	assert.Equal(t, user.Name, createdUser.Name)
	assert.Equal(t, user.Email, createdUser.Email)
}

func TestGetAllUsers(t *testing.T) {
	db := setupTestDB()
	repo := repository.NewUserRepositoryImpl(db)

	var initialCount int64
	db.Model(&model.User{}).Count(&initialCount)

	repo.CreateUser(model.User{Name: "User 1", Email: "user1@example.com"})
	repo.CreateUser(model.User{Name: "User 2", Email: "user2@example.com"})

	users, err := repo.GetAllUsers()

	assert.NoError(t, err)
	assert.Equal(t, int(initialCount)+2, len(users))
}

func TestGetUserByIDRepo(t *testing.T) {
	db := setupTestDB()
	repo := repository.NewUserRepositoryImpl(db)

	user := model.User{Name: "Test User", Email: "test@example.com"}
	createdUser, _ := repo.CreateUser(user)

	fetchedUser, err := repo.GetUserByID(int(createdUser.ID))

	assert.NoError(t, err)
	assert.Equal(t, createdUser.Name, fetchedUser.Name)
	assert.Equal(t, createdUser.Email, fetchedUser.Email)
}

func TestUpdateUserRepo(t *testing.T) {
	db := setupTestDB()
	repo := repository.NewUserRepositoryImpl(db)

	user := model.User{Name: "Test User", Email: "test@example.com"}
	createdUser, _ := repo.CreateUser(user)

	createdUser.Name = "Updated User"
	createdUser.Email = "updated@example.com"
	err := repo.UpdateUser(createdUser)

	assert.NoError(t, err)

	updatedUser, _ := repo.GetUserByID(int(createdUser.ID))
	assert.Equal(t, "Updated User", updatedUser.Name)
	assert.Equal(t, "updated@example.com", updatedUser.Email)
}

func TestDeleteUserRepo(t *testing.T) {
	db := setupTestDB()
	repo := repository.NewUserRepositoryImpl(db)

	user := model.User{Name: "Test User", Email: "test@example.com"}
	createdUser, _ := repo.CreateUser(user)

	err := repo.DeleteUser(int(createdUser.ID))
	assert.NoError(t, err)

	_, err = repo.GetUserByID(int(createdUser.ID))
	assert.Error(t, err)
}
