package handler

import (
	"auction-service/internal/model"
	"auction-service/internal/repository"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"gorm.io/gorm"

	"github.com/gorilla/mux"
)

var auctionRepository *repository.AuctionRepository

func InitAuctionRepository(db *gorm.DB) {
	auctionRepository = repository.NewAuctionRepository(db)
}

// GetAllAuctions handles the request to get all auctions.
func GetAllAuctions(w http.ResponseWriter, r *http.Request) {
	// Logic to fetch all auctions from the database
	auctions, err := auctionRepository.GetAllAuctions()
	if err != nil {
		log.Printf("Error fetching auctions: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Write JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(auctions)
}

// GetAuctionByID handles the request to get an auction by its ID.
func GetAuctionByID(w http.ResponseWriter, r *http.Request) {
	// Get auction ID from request parameters
	vars := mux.Vars(r)
	auctionID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid auction ID", http.StatusBadRequest)
		return
	}

	// Logic to fetch auction by ID from the database
	auction, err := auctionRepository.GetAuctionByID(auctionID)
	if err != nil {
		log.Printf("Error fetching auction by ID: %v", err)
		http.Error(w, "Auction not found", http.StatusNotFound)
		return
	}

	// Write JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(auction)
}

// CreateAuction handles the request to create a new auction.
func CreateAuction(w http.ResponseWriter, r *http.Request) {
	// Decode JSON request body into Auction struct
	var newAuction model.Auction
	if err := json.NewDecoder(r.Body).Decode(&newAuction); err != nil {
		http.Error(w, "Invalid JSON request body", http.StatusBadRequest)
		return
	}

	// Logic to create a new auction in the database
	createdAuction, err := auctionRepository.CreateAuction(newAuction)
	if err != nil {
		log.Printf("Error creating auction: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Write JSON response with the newly created auction
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdAuction)
}

// UpdateAuction handles the request to update an existing auction.
func UpdateAuction(w http.ResponseWriter, r *http.Request) {
	// Get auction ID from request parameters
	vars := mux.Vars(r)
	auctionID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid auction ID", http.StatusBadRequest)
		return
	}

	// Decode JSON request body into Auction struct
	var updatedAuction model.Auction
	if err := json.NewDecoder(r.Body).Decode(&updatedAuction); err != nil {
		http.Error(w, "Invalid JSON request body", http.StatusBadRequest)
		return
	}

	// Logic to update auction in the database
	updatedAuction.ID = uint(auctionID) // Ensure ID matches the one in the URL
	if err := auctionRepository.UpdateAuction(updatedAuction); err != nil {
		log.Printf("Error updating auction: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Write JSON response with the updated auction
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedAuction)
}

// DeleteAuction handles the request to delete an existing auction.
func DeleteAuction(w http.ResponseWriter, r *http.Request) {
	// Get auction ID from request parameters
	vars := mux.Vars(r)
	auctionID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid auction ID", http.StatusBadRequest)
		return
	}

	// Logic to delete auction from the database
	if err := auctionRepository.DeleteAuction(auctionID); err != nil {
		log.Printf("Error deleting auction: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Write JSON response
	w.WriteHeader(http.StatusNoContent)
	fmt.Fprintf(w, "Auction deleted successfully")
}
