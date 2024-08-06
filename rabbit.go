package rabbit

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
