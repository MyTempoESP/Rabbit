package rabbit

import (
	"fmt"
	"os"
)

/* env vars here */
var RABBITMQ_PASS = "RABBITMQ_PASS"
var RABBITMQ_USER = "RABBITMQ_USER"
var RABBITMQ_HOST = "RABBITMQ_HOST"
var RABBITMQ_PORT = "RABBITMQ_PORT"

var UrlFormat = "amqp://%s@%s:%s/"

var ErrNoRabbitKey = fmt.Errorf("RABBITMQ_PASS env not found")
var ErrNoRabbitUser = fmt.Errorf("RABBITMQ_USER env not found")
var ErrNoRabbitHost = fmt.Errorf("RABBITMQ_HOST env not found")

func getRabbitAuth() (user string, pass string, err error) {

	pass, err = os.Getenv(RABBITMQ_PASS), nil

	if pass == "" {
		err = ErrNoRabbitKey
		return
	}

	user = os.Getenv(RABBITMQ_USER)

	if user == "" {
		err = ErrNoRabbitUser

		/* fallthrough */
	}

	return
}

func getRabbitServer() (host string, port string, err error) {

	err = nil

	host = os.Getenv(RABBITMQ_HOST)
	port = os.Getenv(RABBITMQ_PORT)

	if host == "" {
		err = ErrNoRabbitHost
		return
	}

	// FIXME: issue warning or something

	if port == "" {
		port = "5672"

		/* fallthrough */
	}

	return
}

func (rabbit *Rabbit) setupAuth() (err error) {

	var user, pass string

	user, pass, err = getRabbitAuth()

	if err != nil {
		return err
	}

	/* expose state */
	rabbit.user = user
	rabbit.pass = pass

	return nil
}

func (rabbit *Rabbit) setupServer() (err error) {

	var host, port string

	host, port, err = getRabbitServer()

	if err != nil {
		return err
	}

	/* expose state */
	rabbit.host = host
	rabbit.port = port

	return nil
}

/* private */
func (rabbit *Rabbit) url() (url string) {
	/*
		This function uses:

			rabbit.user
			rabbit.pass

			rabbit.host
			rabbit.port

		to generate an amqp url.

		It is not exposed, as it can reveal
		the user password in the environment
		file.
	*/

	var auth string

	/*
		setup a url in the form:

		amqp://<user>:<password>@<host>:<port>
	*/

	auth = fmt.Sprintf("%s:%s", rabbit.user, rabbit.pass)
	url = fmt.Sprintf(UrlFormat, auth, rabbit.host, rabbit.port)

	return
}

func (rabbit *Rabbit) LogUrl() (url string) {
	/*
		Log- variation of Url function,
		this function doesn't expose the
		user password.
	*/

	/*
		This function uses:

			rabbit.user

			rabbit.host
			rabbit.port

		to generate an amqp url suitable for
		logging. (Omitting password)
	*/

	var auth string

	/*
		setup a url in the form:

		amqp://<user>:******@<host>:<port>
	*/

	auth = fmt.Sprintf("%s:******", rabbit.user)
	url = fmt.Sprintf("amqp://%s@%s:%s/", auth, rabbit.host, rabbit.port)

	return
}
