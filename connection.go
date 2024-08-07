package rabbit

import (
	"fmt"

	backoff "github.com/cenkalti/backoff"
	amqp "github.com/rabbitmq/amqp091-go"
)

func (rabbit *Rabbit) Connect() (err error) {

	var url string
	var conn *amqp.Connection /* Connection pointer */

	url = rabbit.url()

	err = backoff.Retry(
		func() (err error) {
			conn, err = amqp.Dial(url)

			return
		},

		backoff.NewExponentialBackOff(),
	)

	if err != nil {
		err = fmt.Errorf("Couldn't connect to RabbitMQ: %s", err)

		return
	}

	rabbit.conn = conn

	return
}

func (rabbit *Rabbit) NotifyClose(c chan *amqp.Error) {
	if rabbit.conn == nil || rabbit.conn.IsClosed() {
		c <- nil

		return
	}

	rabbit.conn.NotifyClose(c)
}
