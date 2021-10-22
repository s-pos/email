package queue

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/streadway/amqp"
)

var Exchange = os.Getenv("RABBIT_EXCHANGE")

type ConsumeFunc func(amqp.Delivery)

type Queue interface {
	// setup queue
	Setup() error

	// Consume will consume the queue and handle the message comming
	Consume(name string) (<-chan amqp.Delivery, error)

	// HandleConsume will handle a message incoming with any worker
	HandleConsume(fn ConsumeFunc) error
}
