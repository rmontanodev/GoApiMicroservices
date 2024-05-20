package handler

import (
	"auction-service/internal/model"
	"auction-service/internal/repository"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"gorm.io/gorm"
)

type AuctionRepository interface {
	GetAllAuctions() ([]model.Auction, error)
	GetAuctionByID(id int) (model.Auction, error)
	CreateAuction(auction model.Auction) (model.Auction, error)
	UpdateAuction(auction model.Auction) error
	DeleteAuction(id int) error
}

type AuctionHandler struct {
	repo AuctionRepository
}

func NewAuctionHandler(repo AuctionRepository) *AuctionHandler {
	return &AuctionHandler{repo: repo}
}

var auctionRepository repository.AuctionRepository

func InitAuctionRepository(db *gorm.DB) repository.AuctionRepository {
	auctionRepository := repository.NewAuctionRepository(db)
	return auctionRepository
}

func (h *AuctionHandler) GetAllAuctions(w http.ResponseWriter, r *http.Request) {
	auctions, err := h.repo.GetAllAuctions()
	if err != nil {
		log.Printf("Error fetching auctions: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(auctions)
}

func (h *AuctionHandler) GetAuctionByID(w http.ResponseWriter, r *http.Request) {
	re := regexp.MustCompile(`^/auctions/(\d+)$`)
	matches := re.FindStringSubmatch(r.URL.Path)
	if len(matches) != 2 {
		http.Error(w, "Invalid auction ID", http.StatusBadRequest)
		return
	}

	auctionID, err := strconv.Atoi(matches[1])
	if err != nil {
		http.Error(w, "Invalid auction ID", http.StatusBadRequest)
		return
	}

	auction, err := h.repo.GetAuctionByID(auctionID)
	if err != nil {
		log.Printf("Error fetching auction by ID: %v", err)
		http.Error(w, "Auction not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(auction)
}

func (h *AuctionHandler) CreateAuction(w http.ResponseWriter, r *http.Request) {
	var newAuction model.Auction
	if err := json.NewDecoder(r.Body).Decode(&newAuction); err != nil {
		http.Error(w, "Invalid JSON request body", http.StatusBadRequest)
		return
	}

	createdAuction, err := h.repo.CreateAuction(newAuction)
	if err != nil {
		log.Printf("Error creating auction: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdAuction)
}

func (h *AuctionHandler) UpdateAuction(w http.ResponseWriter, r *http.Request) {
	re := regexp.MustCompile(`^/auctions/update/(\d+)$`)
	matches := re.FindStringSubmatch(r.URL.Path)
	if len(matches) != 2 {
		http.Error(w, "Invalid auction ID", http.StatusBadRequest)
		return
	}

	auctionID, err := strconv.Atoi(matches[1])
	if err != nil {
		http.Error(w, "Invalid auction ID", http.StatusBadRequest)
		return
	}

	var updatedAuction model.Auction
	if err := json.NewDecoder(r.Body).Decode(&updatedAuction); err != nil {
		http.Error(w, "Invalid JSON request body", http.StatusBadRequest)
		return
	}

	updatedAuction.ID = uint(auctionID)
	if err := h.repo.UpdateAuction(updatedAuction); err != nil {
		log.Printf("Error updating auction: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedAuction)
}

func (h *AuctionHandler) DeleteAuction(w http.ResponseWriter, r *http.Request) {
	re := regexp.MustCompile(`^/auctions/delete/(\d+)$`)
	matches := re.FindStringSubmatch(r.URL.Path)
	if len(matches) != 2 {
		http.Error(w, "Invalid auction ID", http.StatusBadRequest)
		return
	}

	auctionID, err := strconv.Atoi(matches[1])
	if err != nil {
		http.Error(w, "Invalid auction ID", http.StatusBadRequest)
		return
	}

	if err := h.repo.DeleteAuction(auctionID); err != nil {
		log.Printf("Error deleting auction: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	fmt.Fprintf(w, "Auction deleted successfully")
}
