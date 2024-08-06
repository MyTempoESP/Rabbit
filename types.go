package rabbit

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitChannel struct {
	channel *amqp.Channel

	/* Channel for notification on close */
	closed chan *amqp.Error
}

type Rabbit /* RabbitMQ connection */ struct {
	host string
	port string

	user string
	pass string

	exchange   string /* current exchange name */
	routingKey string /*select Queue to which msg will b sent*/

	conn *amqp.Connection /* current open connection (Connection pointer) */

	channel RabbitChannel
}
