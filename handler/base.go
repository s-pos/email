package handler

import "github.com/streadway/amqp"

type BaseHandler interface {
	Handler(msg amqp.Delivery)
}
