package main

import (
	"log"
	"net/http"
	"user-service/internal/config"
	"user-service/internal/handler"
	"user-service/internal/model"
	"user-service/internal/repository"
	"user-service/rabbitmq"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	var publisher *rabbitmq.Publisher

	cfg := config.LoadConfig() // Get DatabaseURL from config

	// Database connection details (use value from config)
	dsn := cfg.DatabaseURL // Use DatabaseURL returned by LoadConfig

	// Connect to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}) // Use postgres.Open with dsn
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = db.AutoMigrate(&model.User{}, &model.User{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Migration successful")

	handler.InitUserRepository(db)

	userHandler := handler.NewUserHandler(&repository.UserRepositoryImpl{}, publisher)
	// Register HTTP endpoints with handler methods
	http.HandleFunc("/users", userHandler.GetAllUsers)
	http.HandleFunc("/users/{id}", userHandler.GetUserByID)
	http.HandleFunc("/users/create", userHandler.CreateUser)
	http.HandleFunc("/users/update/{id}", userHandler.UpdateUser)
	http.HandleFunc("/users/delete/{id}", userHandler.DeleteUser)

	log.Printf("User Service running on port %s", cfg.ServerPort)
	log.Fatal(http.ListenAndServe(":"+cfg.ServerPort, nil))
}
