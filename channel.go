package rabbit

import (
	"time"

	backoff "github.com/cenkalti/backoff"
	amqp "github.com/rabbitmq/amqp091-go"
)

// TODO: enable creating multiple channels
// at once

func (rabbit *Rabbit) Channel() (channel *amqp.Channel, err error) {

	channel = rabbit.channel

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

func (rabbit *Rabbit) EnsureChannel() (channel *amqp.Channel, err error) {

	/* XXX: To back a back-off */

	exp := backoff.NewExponentialBackOff()
	exp.MaxElapsedTime = 20 * time.Second

	err = backoff.Retry(
		func() (err error) {
			channel, err = rabbit.Channel()

			if err != nil {
				rabbit.Reconnect()
			}

			return
		},

		exp,
	)

	return
}
