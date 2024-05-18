package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"user-service/internal/model"
	"user-service/internal/repository"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var userRepository *repository.UserRepository

func InitUserRepository(db *gorm.DB) {
	userRepository = repository.NewUserRepository(db)
}

// GetAllUsers handles the request to get all users.
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	// Logic to fetch all users from the database
	users, err := userRepository.GetAllUsers()
	if err != nil {
		log.Printf("Error fetching users: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Write JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// GetUserByID handles the request to get a user by its ID.
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	// Get user ID from request parameters
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Logic to fetch user by ID from the database
	user, err := userRepository.GetUserByID(userID)
	if err != nil {
		log.Printf("Error fetching user by ID: %v", err)
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Write JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// CreateUser handles the request to create a new user.
func CreateUser(w http.ResponseWriter, r *http.Request) {
	// Decode JSON request body into User struct
	var newUser model.User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "Invalid JSON request body", http.StatusBadRequest)
		return
	}

	// Logic to create a new user in the database
	createdUser, err := userRepository.CreateUser(newUser)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Write JSON response with the newly created user
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser)
}

// UpdateUser handles the request to update an existing user.
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Get user ID from request parameters
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Decode JSON request body into User struct
	var updatedUser model.User
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		http.Error(w, "Invalid JSON request body", http.StatusBadRequest)
		return
	}

	// Logic to update user in the database
	updatedUser.ID = uint(userID) // Ensure ID matches the one in the URL
	if err := userRepository.UpdateUser(updatedUser); err != nil {
		log.Printf("Error updating user: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Write JSON response with the updated user
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedUser)
}

// DeleteUser handles the request to delete an existing user.
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Get user ID from request parameters
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Logic to delete user from the database
	if err := userRepository.DeleteUser(userID); err != nil {
		log.Printf("Error deleting user: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Write JSON response
	w.WriteHeader(http.StatusNoContent)
	fmt.Fprintf(w, "User deleted successfully")
}
