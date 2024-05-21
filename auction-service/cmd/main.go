package main

import (
	"auction-service/internal/config"
	"auction-service/internal/handler"
	"auction-service/internal/model"
	"auction-service/internal/repository"
	"auction-service/rabbitmq"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type RabbitMQConfig struct {
	RabbitMQURL string `json:"rabbitmq_url"`
	QueueName   string `json:"queue_name"`
}

func main() {
	cfg := config.LoadConfig() // Get DatabaseURL from config

	// Database connection details (use value from config)
	dsn := cfg.DatabaseURL // Use DatabaseURL returned by LoadConfig

	// Connect to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}) // Use postgres.Open with dsn
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Migrar el esquema de User
	err = db.AutoMigrate(&model.User{}, &model.Auction{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Migration successful")
	repo := repository.NewAuctionRepository(db)

	consumer, err := rabbitmq.NewConsumer(config.LoadConfig().RabbitMQURL, config.LoadConfig().QUEUE_USER_CREATED, repo)
	if err != nil {
		log.Fatalf("Failed to create RabbitMQ consumer: %v", err)
	}

	consumer.StartConsuming()

	// Initialize auction repository with database connection
	handler.InitAuctionRepository(db)

	// Create an AuctionHandler instance
	var auctionHandler handler.AuctionHandler

	// Register HTTP endpoints with handler methods
	http.HandleFunc("/auctions", auctionHandler.GetAllAuctions)
	http.HandleFunc("/auctions/{id}", auctionHandler.GetAuctionByID)
	http.HandleFunc("/auctions/create", auctionHandler.CreateAuction)
	http.HandleFunc("/auctions/update/{id}", auctionHandler.UpdateAuction)
	http.HandleFunc("/auctions/delete/{id}", auctionHandler.DeleteAuction)

	log.Printf("Auction Service running on port %s", cfg.ServerPort)
	log.Fatal(http.ListenAndServe(":"+cfg.ServerPort, nil))
}
