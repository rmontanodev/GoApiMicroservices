package main

import (
	"log"
	"net/http"
	"user-service/internal/config"
	"user-service/internal/handler"
)

func main() {
	cfg := config.LoadConfig()

	http.HandleFunc("/users", handler.GetAllUsers)
	http.HandleFunc("/users/:id", handler.GetUserByID)
	http.HandleFunc("/users/create", handler.CreateUser)
	http.HandleFunc("/users/:id", handler.UpdateUser)
	http.HandleFunc("/users/:id", handler.DeleteUser)

	log.Printf("User Service running on port %s", cfg.ServerPort)
	log.Fatal(http.ListenAndServe(":"+cfg.ServerPort, nil))
}
