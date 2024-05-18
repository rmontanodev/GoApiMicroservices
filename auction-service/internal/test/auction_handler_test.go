package test

import (
	"auction/auction-service/internal/handler"
	"auction/auction-service/internal/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// Mock de AuctionService
type MockAuctionService struct{}

func (m *MockAuctionService) CreateAuction(item string, userID uint) error {
	return nil
}

func (m *MockAuctionService) GetAuctionByID(id uint) (*model.Auction, error) {
	return &model.Auction{Model: gorm.Model{ID: id}, Item: "Test Item", UserID: 1}, nil
}

func TestCreateAuction(t *testing.T) {
	r := gin.Default()
	auctionService := &MockAuctionService{}
	auctionHandler := handler.NewAuctionHandler(auctionService)
	r.POST("/auctions", auctionHandler.CreateAuction)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/auctions", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetAuction(t *testing.T) {
	r := gin.Default()
	auctionService := &MockAuctionService{}
	auctionHandler := handler.NewAuctionHandler(auctionService)
	r.GET("/auctions/:id", auctionHandler.GetAuction)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/auctions/1", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Test Item")
}
