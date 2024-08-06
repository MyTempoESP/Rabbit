package rabbit

/*
example:

	...

	import (
		rabbit "..."
	)

	func main() {
		var r rabbit.Rabbit

		err := r.Setup()
		// r.TopicQueue("api_topic", "api_data")

		err := r.Topic("api_topic")
		if err != nil {
			return 1
		}

		err = r.QueueDeclare("api.data", true)
		if err != nil {
			return 1
		}

		//... fetch data from api ...

		r.SendMessage(apiData)
	}
*/

import (
	"fmt"
)

func (rabbit *Rabbit) Setup() (err error) {
	err = rabbit.setupServer()

	if err != nil {
		return fmt.Errorf("Failed to setup server: %s", err)
	}

	err = rabbit.setupAuth()

	if err != nil {
		return fmt.Errorf("Failed to setup authentication: %s", err)
	}

	err = rabbit.Connect()

	if err != nil {
		return fmt.Errorf("Failed to connect to RabbitMQ: %s", err)
	}

	return
}
