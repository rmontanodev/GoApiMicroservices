package rabbitmq

import (
    "log"
    "github.com/streadway/amqp"
)

func NewRabbitMQConnection(url string) (*amqp.Connection, error) {
    conn, err := amqp.Dial(url)
    if err != nil {
        return nil, err
    }
    return conn, nil
}

func FailOnError(err error, msg string) {
    if err != nil {
        log.Fatalf("%s: %s", msg, err)
    }
}
