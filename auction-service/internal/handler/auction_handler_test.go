package handler_test

import (
	"auction-service/internal/handler"
	"auction-service/internal/model"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockAuctionRepository is a mock implementation of the AuctionRepository interface
type MockAuctionRepository struct {
	mock.Mock
}

func (m *MockAuctionRepository) GetAllAuctions() ([]model.Auction, error) {
	args := m.Called()
	return args.Get(0).([]model.Auction), args.Error(1)
}

func (m *MockAuctionRepository) GetAuctionByID(id int) (model.Auction, error) {
	args := m.Called(id)
	return args.Get(0).(model.Auction), args.Error(1)
}

func (m *MockAuctionRepository) CreateAuction(auction model.Auction) (model.Auction, error) {
	args := m.Called(auction)
	return args.Get(0).(model.Auction), args.Error(1)
}

func (m *MockAuctionRepository) UpdateAuction(auction model.Auction) error {
	args := m.Called(auction)
	return args.Error(0)
}

func (m *MockAuctionRepository) DeleteAuction(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestCreateAuction(t *testing.T) {
	mockRepo := new(MockAuctionRepository)
	auction := model.Auction{Item: "Test Item", UserID: 1}
	mockRepo.On("CreateAuction", auction).Return(auction, nil)

	auctionHandler := handler.NewAuctionHandler(mockRepo)

	reqBody, _ := json.Marshal(auction)
	req, err := http.NewRequest("POST", "/auctions", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(auctionHandler.CreateAuction)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	mockRepo.AssertExpectations(t)
}

func TestGetAuctionByID(t *testing.T) {
	mockRepo := new(MockAuctionRepository)

	auction := model.Auction{Item: "Test Item", UserID: 1}

	mockRepo.On("GetAuctionByID", 1).Return(auction, nil)

	auctionHandler := handler.NewAuctionHandler(mockRepo)

	req, err := http.NewRequest("GET", "/auctions/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(auctionHandler.GetAuctionByID)
	handler.ServeHTTP(rr, req)

	// Imprimir el cuerpo de la respuesta para la depuración
	t.Logf("Response body: %s", rr.Body.String())

	assert.Equal(t, http.StatusOK, rr.Code)

	var returnedAuction model.Auction
	err = json.Unmarshal(rr.Body.Bytes(), &returnedAuction)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, auction.Item, returnedAuction.Item)
	assert.Equal(t, auction.UserID, returnedAuction.UserID)

	mockRepo.AssertExpectations(t)
}
