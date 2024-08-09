package rabbit

import (
// amqp "github.com/rabbitmq/amqp091-go"
)

func (rabbit *Rabbit) BindQueues(bindings []Binding) (err error) {

	channel, err := rabbit.Conn.Channel()

	if err != nil {
		return
	}

	defer channel.Close()

	for _, b := range bindings {
		_, err = channel.QueueDeclare(b.Queue, b.Durable, b.AutoDelete, b.Exclusive, false, nil)

		if err != nil {
			return
		}

		err = channel.QueueBind(b.Queue, b.Key, b.Exchange, false, nil)

		if err != nil {
			return
		}
	}

	return
}
