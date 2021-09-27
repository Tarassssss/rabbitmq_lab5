package main

import (
	"fmt"
	"github.com/streadway/amqp"
)

func main() {
	conn, _ := amqp.Dial("amqp://guest:guest@localhost:5672/")
	ch, _ := conn.Channel()
	q, _ := ch.QueueDeclare("limitedQueue5", false, false, false, false, amqp.Table{"x-max-length": 10, "x-overflow": "reject-publish"})
	msgs, _ := ch.Consume(q.Name, "", true, false, false, false, nil)

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Println("Received a message:", string(d.Body))
		}
	}()
	<-forever
}
