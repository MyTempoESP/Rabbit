package rabbit

import (
	"fmt"

	backoff "github.com/cenkalti/backoff"
	amqp "github.com/rabbitmq/amqp091-go"
)

func (rabbit *Rabbit) Connect() (err error) {

	var url string
	var Conn *amqp.Connection /* Connection pointer */

	url = rabbit.url()

	err = backoff.Retry(
		func() (err error) {
			Conn, err = amqp.Dial(url)

			return
		},

		backoff.NewExponentialBackOff(),
	)

	if err != nil {
		err = fmt.Errorf("Couldn't Connect to RabbitMQ: %s", err)

		return
	}

	rabbit.Conn = Conn

	return
}

func (rabbit *Rabbit) NotifyClose(c chan *amqp.Error) {
	if rabbit.Conn == nil || rabbit.Conn.IsClosed() {
		close(c)

		return
	}

	rabbit.Conn.NotifyClose(c)
}
