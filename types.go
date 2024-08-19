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

func NewBinding(queue string, key string, exchange string, opts ...bool) (bound Binding) {

	bound.Exchange = exchange
	bound.Queue = queue
	bound.Key = key

	/* set 'Durable', 'AutoDelete' and 'Exclusive' accordingly */
	switch len(opts) {
	case 3:
		/*
			falls through if
			the length of
			arguments is 3,
			thus setting the
			rest of the fields
			in the struct.
		*/
		bound.Exclusive = opts[2]
		fallthrough
	case 2:
		bound.AutoDelete = opts[1]
		fallthrough
	case 1:
		bound.Durable = opts[0]
	}

	return
}
