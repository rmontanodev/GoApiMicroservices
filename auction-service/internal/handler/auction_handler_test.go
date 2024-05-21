package handler_test

import (
	"auction-service/internal/handler"
	"auction-service/internal/model"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
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
func TestUpdateAuction(t *testing.T) {
	mockRepo := new(MockAuctionRepository)
	initialAuction := model.Auction{ID: 1, Item: "Initial Item", UserID: 1}
	updatedAuction := model.Auction{ID: 1, Item: "Updated Item", UserID: 1}

	mockRepo.On("CreateAuction", initialAuction).Return(initialAuction, nil)

	mockRepo.On("UpdateAuction", updatedAuction).Return(nil)

	auctionHandler := handler.NewAuctionHandler(mockRepo)

	createdAuction, err := mockRepo.CreateAuction(initialAuction)
	if err != nil {
		t.Fatal(err)
	}

	reqBody, _ := json.Marshal(updatedAuction)
	req, err := http.NewRequest("PUT", "/auctions/update/1", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(auctionHandler.UpdateAuction)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	assert.NotEqual(t, updatedAuction.Item, createdAuction.Item)
	assert.Equal(t, updatedAuction.UserID, createdAuction.UserID)

	mockRepo.AssertExpectations(t)
}

func TestGetAllAuctions(t *testing.T) {
	mockRepo := new(MockAuctionRepository)
	auctions := []model.Auction{
		{Item: "Test Item 1", UserID: 1},
		{Item: "Test Item 2", UserID: 2},
	}
	mockRepo.On("GetAllAuctions").Return(auctions, nil)

	auctionHandler := handler.NewAuctionHandler(mockRepo)

	req, err := http.NewRequest("GET", "/auctions", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(auctionHandler.GetAllAuctions)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var returnedAuctions []model.Auction
	err = json.Unmarshal(rr.Body.Bytes(), &returnedAuctions)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, auctions, returnedAuctions)
	mockRepo.AssertExpectations(t)
}

func TestDeleteAuction(t *testing.T) {
	mockRepo := new(MockAuctionRepository)
	auction := model.Auction{ID: 1, Item: "Test Item", UserID: 1}

	// Mock the CreateAuction method
	mockRepo.On("CreateAuction", auction).Return(auction, nil)

	auctionHandler := handler.NewAuctionHandler(mockRepo)

	// Step 1: Create an auction
	auctionJSON, _ := json.Marshal(auction)
	reqCreate, err := http.NewRequest("POST", "/auctions", bytes.NewBuffer(auctionJSON))
	if err != nil {
		t.Fatal(err)
	}
	reqCreate.Header.Set("Content-Type", "application/json")

	rrCreate := httptest.NewRecorder()
	createHandler := http.HandlerFunc(auctionHandler.CreateAuction)
	createHandler.ServeHTTP(rrCreate, reqCreate)

	assert.Equal(t, http.StatusCreated, rrCreate.Code)

	var createdAuction model.Auction
	err = json.NewDecoder(rrCreate.Body).Decode(&createdAuction)
	if err != nil {
		t.Fatal(err)
	}

	// Mock the DeleteAuction method
	mockRepo.On("DeleteAuction", createdAuction.ID).Return(nil)

	// Step 2: Delete the created auction
	reqDelete, err := http.NewRequest("DELETE", "/auctions/delete/"+strconv.Itoa(int(createdAuction.ID)), nil)
	if err != nil {
		t.Fatal(err)
	}

	rrDelete := httptest.NewRecorder()
	deleteHandler := http.HandlerFunc(auctionHandler.DeleteAuction)
	deleteHandler.ServeHTTP(rrDelete, reqDelete)

	assert.Equal(t, http.StatusNoContent, rrDelete.Code)

	mockRepo.AssertExpectations(t)
}
