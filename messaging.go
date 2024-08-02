package rabbit

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func (rabbit *Rabbit) SendTo(msg string, exchangeName string, routingKey string, timeout_s int) (err error) {

	var channel *amqp.Channel

	log.Println("Trying to send message through RabbitMQ")

	timeout_duration := time.Duration(timeout_s)

	ctx, cancel := context.WithTimeout(context.Background(), timeout_duration*time.Second)
	defer cancel()

	channel, err = rabbit.Channel()

	if err != nil {
		return /* up to the caller to re-send */
	}

	defer channel.Close()

	log.Printf("Publishing message to exchange %s, queue(s): %s", rabbit.exchange, rabbit.routingKey)

	err = channel.PublishWithContext(ctx,
		rabbit.exchange, /* it's as easy as this */
		rabbit.routingKey,
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain", /* "application/json", */
			Body:        []byte(msg),

			/*	NOTE:
				for Persistant messaging, uncomment
				the following line:
			*/
			//DeliveryMode: 2, /* default: 0 (transient) */
		})

	return
}

func (rabbit *Rabbit) SendMessage(msg string, timeout_s int) (err error) {

	err = rabbit.SendTo(msg, rabbit.exchange, rabbit.routingKey, timeout_s)

	return
}
