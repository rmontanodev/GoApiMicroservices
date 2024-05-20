// internal/repository/auction_repository.go
package repository

import (
	"auction-service/internal/model"
)

// AuctionRepository defines the methods that any repository implementation must have.
type AuctionRepository interface {
	GetAllAuctions() ([]model.Auction, error)
	GetAuctionByID(id int) (model.Auction, error)
	CreateAuction(auction model.Auction) (model.Auction, error)
	UpdateAuction(auction model.Auction) error
	DeleteAuction(id int) error
}
