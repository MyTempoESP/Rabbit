package rabbit

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func (rabbit *Rabbit) Channel() (channel *amqp.Channel, err error) {

	channel, err = rabbit.conn.Channel()

	return
}
