package repository

import (
	"auction-service/internal/model"

	"gorm.io/gorm"
)

// AuctionRepository handles database operations related to auctions.
type AuctionRepository struct {
	db *gorm.DB
}

// NewAuctionRepository creates a new instance of AuctionRepository.
func NewAuctionRepository(db *gorm.DB) *AuctionRepository {
	return &AuctionRepository{db}
}

// GetAllAuctions returns all auctions from the database.
func (ar *AuctionRepository) GetAllAuctions() ([]model.Auction, error) {
	var auctions []model.Auction
	err := ar.db.Find(&auctions).Error
	return auctions, err
}

// GetAuctionByID returns an auction by its ID from the database.
func (ar *AuctionRepository) GetAuctionByID(id int) (model.Auction, error) {
	var auction model.Auction
	err := ar.db.First(&auction, id).Error
	return auction, err
}

// CreateAuction creates a new auction in the database.
func (ar *AuctionRepository) CreateAuction(auction model.Auction) (model.Auction, error) {
	err := ar.db.Create(&auction).Error
	return auction, err
}

// UpdateAuction updates an existing auction in the database.
func (ar *AuctionRepository) UpdateAuction(updatedAuction model.Auction) error {
	err := ar.db.Save(&updatedAuction).Error
	return err
}

// DeleteAuction deletes an existing auction from the database by its ID.
func (ar *AuctionRepository) DeleteAuction(id int) error {
	err := ar.db.Delete(&model.Auction{}, id).Error
	return err
}
