package rabbitmq

import (
	"github.com/streadway/amqp"
)

type Publisher struct {
	Channel *amqp.Channel
	Queue   amqp.Queue
}

func NewPublisher(url, queueName string) (*Publisher, error) {
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

	return &Publisher{
		Channel: ch,
		Queue:   q,
	}, nil
}

func (p *Publisher) Publish(message []byte) error {
	return p.Channel.Publish(
		"",
		p.Queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		},
	)
}
