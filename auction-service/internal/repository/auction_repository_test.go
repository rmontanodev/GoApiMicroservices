package repository_test

import (
	"auction-service/internal/config"
	"auction-service/internal/model"
	"auction-service/internal/repository"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	cfg := config.LoadConfig() // Get DatabaseURL from config

	// Database connection details (use value from config)
	dsn := cfg.DatabaseURL // Use DatabaseURL returned by LoadConfig

	// Connect to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}) // Use postgres.Open with dsn
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	return db
}

func TestCreateAuctionRepo(t *testing.T) {
	db := setupTestDB()
	repo := repository.NewAuctionRepository(db)

	auction := model.Auction{Item: "Test Item", UserID: 1}
	createdAuction, err := repo.CreateAuction(auction)

	assert.NoError(t, err)
	assert.Equal(t, auction.Item, createdAuction.Item)
	assert.Equal(t, auction.UserID, createdAuction.UserID)
}

func TestGetAllAuctions(t *testing.T) {
	db := setupTestDB()
	repo := repository.NewAuctionRepository(db)

	var initialCount int64
	db.Model(&model.Auction{}).Count(&initialCount)

	repo.CreateAuction(model.Auction{Item: "Test Item 1", UserID: 1})
	repo.CreateAuction(model.Auction{Item: "Test Item 2", UserID: 2})

	auctions, err := repo.GetAllAuctions()

	assert.NoError(t, err)
	assert.Equal(t, int(initialCount)+2, len(auctions))
}

func TestGetAuctionByIDRepo(t *testing.T) {
	db := setupTestDB()
	repo := repository.NewAuctionRepository(db)

	auction := model.Auction{Item: "Test Item", UserID: 1}
	createdAuction, _ := repo.CreateAuction(auction)

	fetchedAuction, err := repo.GetAuctionByID(int(createdAuction.ID))

	assert.NoError(t, err)
	assert.Equal(t, createdAuction.Item, fetchedAuction.Item)
	assert.Equal(t, createdAuction.UserID, fetchedAuction.UserID)
}

func TestUpdateAuctionRepo(t *testing.T) {
	db := setupTestDB()
	repo := repository.NewAuctionRepository(db)

	auction := model.Auction{Item: "Test Item", UserID: 1}
	createdAuction, _ := repo.CreateAuction(auction)

	createdAuction.Item = "Updated Item"
	err := repo.UpdateAuction(createdAuction)

	assert.NoError(t, err)

	updatedAuction, _ := repo.GetAuctionByID(int(createdAuction.ID))
	assert.Equal(t, "Updated Item", updatedAuction.Item)
}

func TestDeleteAuctionRepo(t *testing.T) {
	db := setupTestDB()
	repo := repository.NewAuctionRepository(db)

	auction := model.Auction{Item: "Test Item", UserID: 1}
	createdAuction, _ := repo.CreateAuction(auction)

	err := repo.DeleteAuction(int(createdAuction.ID))
	assert.NoError(t, err)

	_, err = repo.GetAuctionByID(int(createdAuction.ID))
	assert.Error(t, err)
}
