package rabbit

import (
	"fmt"
)

var ErrNoRabbitKey = fmt.Errorf("RABBITMQ_PASS env not found")
var ErrNoRabbitUser = fmt.Errorf("RABBITMQ_USER env not found")
var ErrNoRabbitHost = fmt.Errorf("RABBITMQ_HOST env not found")
