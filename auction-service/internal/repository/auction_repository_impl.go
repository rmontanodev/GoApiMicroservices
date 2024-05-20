// internal/repository/auction_repository_impl.go
package repository

import (
	"auction-service/internal/model"

	"gorm.io/gorm"
)

// AuctionRepositoryImpl handles database operations related to auctions.
type AuctionRepositoryImpl struct {
	db *gorm.DB
}

// NewAuctionRepository creates a new instance of AuctionRepository.
func NewAuctionRepository(db *gorm.DB) *AuctionRepositoryImpl {
	return &AuctionRepositoryImpl{db}
}

// Ensure AuctionRepositoryImpl implements AuctionRepository
var _ AuctionRepository = (*AuctionRepositoryImpl)(nil)

// GetAllAuctions returns all auctions from the database.
func (ar *AuctionRepositoryImpl) GetAllAuctions() ([]model.Auction, error) {
	var auctions []model.Auction
	err := ar.db.Find(&auctions).Error
	return auctions, err
}

// GetAuctionByID returns an auction by its ID from the database.
func (ar *AuctionRepositoryImpl) GetAuctionByID(id int) (model.Auction, error) {
	var auction model.Auction
	err := ar.db.First(&auction, id).Error
	return auction, err
}

// CreateAuction creates a new auction in the database.
func (ar *AuctionRepositoryImpl) CreateAuction(auction model.Auction) (model.Auction, error) {
	err := ar.db.Create(&auction).Error
	return auction, err
}

// UpdateAuction updates an existing auction in the database.
func (ar *AuctionRepositoryImpl) UpdateAuction(updatedAuction model.Auction) error {
	err := ar.db.Save(&updatedAuction).Error
	return err
}

// DeleteAuction deletes an existing auction from the database by its ID.
func (ar *AuctionRepositoryImpl) DeleteAuction(id int) error {
	err := ar.db.Delete(&model.Auction{}, id).Error
	return err
}
