package rabbit

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

// TODO: enable creating multiple channels
// at once

func (rabbit *Rabbit) Channel() (channel *amqp.Channel, err error) {

	if rabbit.channel != nil && !rabbit.channel.IsClosed() {
		return
	}

	channel, err = rabbit.Conn.Channel()

	if err != nil {
		return
	}

	rabbit.channel = channel

	return
}
