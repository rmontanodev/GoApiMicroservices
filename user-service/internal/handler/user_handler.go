package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"user-service/internal/config"
	"user-service/internal/model"
	"user-service/internal/repository"
	"user-service/rabbitmq"

	"gorm.io/gorm"
)

type UserHandler struct {
	repo      repository.UserRepository
	publisher *rabbitmq.Publisher
}

func NewUserHandler(repo repository.UserRepository, publisher *rabbitmq.Publisher) *UserHandler {
	return &UserHandler{repo: repo, publisher: publisher}
}

func InitUserRepository(db *gorm.DB) repository.UserRepository {
	return repository.NewUserRepositoryImpl(db)
}

type Config struct {
	RabbitMQURL string `json:"rabbitmq_url"`
	QueueName   string `json:"queue_name"`
}

// GetAllUsers handles the request to get all users.
func (uh *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := uh.repo.GetAllUsers()
	if err != nil {
		log.Printf("Error fetching users: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// GetUserByID handles the request to get a user by its ID.
func (uh *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	// Expresión regular para extraer el ID del usuario de la URL
	re := regexp.MustCompile(`/users/(\d+)$`)
	matches := re.FindStringSubmatch(r.URL.Path)
	if len(matches) != 2 {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Convertir el ID del usuario a entero
	userID, err := strconv.Atoi(matches[1])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Obtener el usuario por su ID desde la base de datos
	user, err := uh.repo.GetUserByID(int(userID))
	if err != nil {
		log.Printf("Error fetching user by ID: %v", err)
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Escribir la respuesta JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// CreateUser handles the request to create a new user.
func (uh *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser model.User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "Invalid JSON request body", http.StatusBadRequest)
		return
	}

	createdUser, err := uh.repo.CreateUser(newUser)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser)

	publisher, err := rabbitmq.NewPublisher(config.LoadConfig().RabbitMQURL, config.LoadConfig().QUEUE_USER_CREATED)
	if err != nil {
		log.Printf("Error creating RabbitMQ publisher: %v", err)
		return
	}

	message := []byte("New user created: " + newUser.Name)
	err = publisher.Publish(message)
	if err != nil {
		log.Printf("Error publishing message to RabbitMQ: %v", err)
		return
	}
}

// UpdateUser handles the request to update an existing user.
func (uh *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	re := regexp.MustCompile(`/users/update/(\d+)$`)
	matches := re.FindStringSubmatch(r.URL.Path)
	if len(matches) != 2 {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(matches[1])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var updatedUser model.User
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		http.Error(w, "Invalid JSON request body", http.StatusBadRequest)
		return
	}

	updatedUser.ID = int(userID)

	if err := uh.repo.UpdateUser(updatedUser); err != nil {
		log.Printf("Error updating user: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedUser)
}

// DeleteUser handles the request to delete an existing user.
func (uh *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	re := regexp.MustCompile(`/users/delete/(\d+)$`)
	matches := re.FindStringSubmatch(r.URL.Path)
	if len(matches) != 2 {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(matches[1])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	if err := uh.repo.DeleteUser(userID); err != nil {
		log.Printf("Error deleting user: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
