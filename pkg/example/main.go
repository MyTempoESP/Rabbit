package main

import (
	//"fmt"
	"log"

	rabbit "github.com/mytempoesp/rabbit"
)

func main() {
	var r rabbit.Rabbit

	err := r.Setup()

	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println("%e\n", err)

	r.CreateExchange(true, "ola_exchange")
	r.CreateQueue(true, "oi_queue")

	r.SendMessage("auixe anu e o ane mi", 5)
}
