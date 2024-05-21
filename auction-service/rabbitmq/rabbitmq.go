package rabbitmq

import (
	"encoding/json"
	"log"

	"auction-service/internal/model"
	"auction-service/internal/repository"

	"github.com/streadway/amqp"
)

type Consumer struct {
	Channel *amqp.Channel
	Queue   amqp.Queue
	Repo    repository.AuctionRepository
}

func NewConsumer(url, queueName string, repo repository.AuctionRepository) (*Consumer, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	q, err := ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &Consumer{
		Channel: ch,
		Queue:   q,
		Repo:    repo,
	}, nil
}

func (c *Consumer) StartConsuming() {
	msgs, err := c.Channel.Consume(
		c.Queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	go func() {
		for d := range msgs {
			var user model.User
			if err := json.Unmarshal(d.Body, &user); err != nil {
				log.Printf("Error decoding message: %v", err)
				continue
			}

			// Process the message (e.g., create a welcome auction for the new user)
			auction := model.Auction{
				Item:   "Welcome Item for " + user.Name,
				UserID: user.ID,
			}
			if _, err := c.Repo.CreateAuction(auction); err != nil {
				log.Printf("Error creating auction: %v", err)
			}
		}
	}()
}
