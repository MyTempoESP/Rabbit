package rabbit

import (
	"fmt"
	"log"

	backoff "github.com/cenkalti/backoff"
	amqp "github.com/rabbitmq/amqp091-go"
)

func (rabbit *Rabbit) Connect() (err error) {

	var url string
	var conn *amqp.Connection /* Connection pointer */

	url = rabbit.url()

	log.Printf("Initiating connection attempt to '%s'", rabbit.LogUrl())

	err = backoff.Retry(
		func() error {
			log.Println("Trying connection to RabbitMQ")

			conn, err = amqp.Dial(url)

			if err != nil {
				log.Printf("Failed to connect to RabbitMQ: %v.\nRetrying...", err)
			}

			return err
		},

		backoff.NewExponentialBackOff(),
	)

	if err != nil {
		err = fmt.Errorf("Couldn't connect to RabbitMQ: %s", err)

		return
	}

	rabbit.conn = conn

	log.Printf("Succesfully connected to RabbitMQ at '%s'", rabbit.LogUrl())

	return
}
