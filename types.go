package rabbit

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type Rabbit /* RabbitMQ connection */ struct {
	host string
	port string

	user string
	pass string

	conn *amqp.Connection /* current open connection (Connection pointer) */
}

type Producer struct {
	rabbit *Rabbit
}

type Consumer struct {
	rabbit *Rabbit
}
