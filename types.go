package rabbit

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type Rabbit /* RabbitMQ connection */ struct {
	host string
	port string

	user string
	pass string

	Conn *amqp.Connection /* current open connection (Connection pointer) */

	channel *amqp.Channel
}

type Binding struct {
	Queue    string
	Key      string
	Exchange string

	/* optional */
	Durable, AutoDelete, Exclusive bool
}
