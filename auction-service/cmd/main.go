package main

import (
	"auction-service/internal/config"
	"auction-service/internal/handler"
	"log"
	"net/http"
)

func main() {
	cfg := config.LoadConfig()

	http.HandleFunc("/auctions", handler.GetAllAuctions)
	http.HandleFunc("/auctions/:id", handler.GetAuctionByID)
	http.HandleFunc("/auctions/create", handler.CreateAuction)
	http.HandleFunc("/auctions/:id", handler.UpdateAuction)
	http.HandleFunc("/auctions/:id", handler.DeleteAuction)

	log.Printf("Auction Service running on port %s", cfg.ServerPort)
	log.Fatal(http.ListenAndServe(":"+cfg.ServerPort, nil))
}
