package service

import (
	"auction/auction-service/internal/model"
	"auction/auction-service/internal/repository"

	"github.com/streadway/amqp"
)

type AuctionService interface {
	CreateAuction(item string, userID uint) error
	GetAuctionByID(id uint) (*model.Auction, error)
	PlaceBid(auctionID uint, bidAmount float64) error
}

type auctionService struct {
	auctionRepository repository.AuctionRepository
	rabbitConn        *amqp.Connection
}

func NewAuctionService(auctionRepository repository.AuctionRepository, rabbitConn *amqp.Connection) AuctionService {
	return &auctionService{
		auctionRepository: auctionRepository,
		rabbitConn:        rabbitConn,
	}
}

func (s *auctionService) CreateAuction(item string, userID uint) error {
	auction := &model.Auction{
		Item:   item,
		UserID: userID,
	}
	return s.auctionRepository.Create(auction)
}

func (s *auctionService) GetAuctionByID(id uint) (*model.Auction, error) {
	return s.auctionRepository.GetByID(id)
}

func (s *auctionService) PlaceBid(auctionID uint, bidAmount float64) error {
	// Aquí es donde utilizarías RabbitMQ para enviar un mensaje
	// al servicio de pujas para procesar la oferta.
	return nil
}
