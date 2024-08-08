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

type Bindings struct {
	exchange string
	queue    string
	key      string
}

type Queue struct {
	Name      string
	isDurable bool

	/* Binding */
	RoutingKey   string
	ExchangeName string

	/* internal queue */
	queue amqp.Queue
}
