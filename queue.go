package rabbit

import ()

/*

	q := rabbit.Queue{
		"", false
	}

	q.Declare(mq)

	q.Bind("",  "", mq)

*/

func (q *Queue) Declare(rabbit *Rabbit) (err error) {
	channel, err := rabbit.Conn.Channel()

	if err != nil {
		return
	}

	defer channel.Close()

	q.queue, err = channel.QueueDeclare(
		q.Name,      // name
		q.isDurable, // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)

	q.Name = q.queue.Name

	return
}

func (rabbit *Rabbit) BindQueues() (err error) {

	bindings := []bind{
		{"page", "alert"},
		{"email", "info"},
		{"firehose", "#"},
	}

	for _, b := range bindings {
		_, err = c.QueueDeclare(b.queue, true, false, false, false, nil)
		if err != nil {
			log.Fatalf("queue.declare: %v", err)
		}

		err = c.QueueBind(b.queue, b.key, "logs", false, nil)
		if err != nil {
			log.Fatalf("queue.bind: %v", err)
		}
	}
}

func (q *Queue) Bind(exchangeName string, routingKey string, rabbit *Rabbit) (err error) {
	channel, err := rabbit.Conn.Channel()

	if err != nil {
		return
	}

	defer channel.Close()

	err = channel.QueueBind(
		q.Name,       // queue name
		routingKey,   // routing key
		exchangeName, // exchange
		false,        // no-wait
		nil,          // arguments
	)

	return
}
