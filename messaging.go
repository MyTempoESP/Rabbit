package rabbit

import (
	"context"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func (rabbit *Rabbit) SendTo(msg string, exchangeName string, routingKey string, timeout_s int) (err error) {

	var channel *amqp.Channel

	timeout_duration := time.Duration(timeout_s)

	ctx, cancel := context.WithTimeout(context.Background(), timeout_duration*time.Second)
	defer cancel()

	channel, err = rabbit.Channel()

	if err != nil {
		return /* up to the caller to re-send */
	}

	defer channel.Close()

	err = channel.PublishWithContext(ctx,
		exchangeName, /* it's as easy as this */
		routingKey,
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain", /* "application/json", */
			Body:        []byte(msg),

			/*
				NOTE:
				for Persistant messaging, uncomment
				the following line:
			*/
			//DeliveryMode: 2, /* default: 0 (transient) */
		})

	return
}
