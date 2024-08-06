package rabbit

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func (rabbit *Rabbit) ExchangeDeclare(chanel *amqp.Channel, exchangeName string, exchangeType string, isDurable bool) (err error) {

	/* NOTE:
	This function opens a Channel.
	*/

	var channel *amqp.Channel

	/*

		Declare a RabbitMQ exchange.
		returns an error if the exchange
		could not be created, otherwise the
		exchange is created and the exchange is
		set accordingly.

		sets the 'rabbit.exchange' field to the
		created exchange name.

	*/

	log.Printf("Declaring exchange %s", exchangeName)

	err = channel.ExchangeDeclare(
		exchangeName,
		exchangeType,
		isDurable,
		false, // auto-deleted
		false, // internal
		false, // no-wait
		nil,   // arguments
	)

	if err != nil {
		return /* don't set the rabbit exchange */
	}

	rabbit.exchange = exchangeName

	return
}

func (rabbit *Rabbit) CreateExchange(isDurable bool, exchangeType string, exchangeNames ...string) (err error) {

	/* NOTE:
	This function opens a Channel.
	*/

	var channel *amqp.Channel

	/*
		Declare one or more RabbitMQ exchanges.
	*/

	channel, err = rabbit.Channel()

	if err != nil {
		return
	}

	defer channel.Close()

	for _, x := range exchangeNames {
		rabbit.ExchangeDeclare(channel, x, exchangeType, isDurable)
	}

	return
}
