package rabbit

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func (rabbit *Rabbit) QueueDeclare(channel *amqp.Channel, isDurable bool, queueName string) (err error) {
	/*
		Declare a RabbitMQ Queue.

		The resulting queue is defined but not returned,
		it should be accessed using a 'routing_key' which
		matches the name of the queue.

		sets the 'rabbit.routing_key' field to the
		created queue name.

		the argument isDurable decides if the queue messages
		will persist after a broker restart

		returns an error status in case the Queue cannot be created.
	*/

	log.Printf("Declaring queue %s", queueName)

	_, err = channel.QueueDeclare(
		queueName,
		isDurable,
		false, // delete when unused (auto-deleted)
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)

	rabbit.routingKey = queueName

	return

}

func (rabbit *Rabbit) CreateQueue(isDurable bool, queueNames ...string) (err error) {

	/* NOTE:
	This function opens a Channel.
	*/

	var channel *amqp.Channel

	/*
		Declare one or more RabbitMQ queues.
	*/

	channel, err = rabbit.Channel()

	if err != nil {
		return
	}

	defer channel.Close()

	for _, q := range queueNames {
		rabbit.QueueDeclare(channel, isDurable, q)
	}

	return
}
