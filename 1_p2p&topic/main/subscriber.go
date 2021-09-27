package main

import (
	"fmt"
	"github.com/streadway/amqp"
)

func main() {
	conn, _ := amqp.Dial("amqp://guest:guest@localhost:5672/")
	ch, _ := conn.Channel()
	_ = ch.ExchangeDeclare("task_subscribe", "fanout", true, false, false, false, nil)
	q, _ := ch.QueueDeclare("", false, false, true, false, nil)
	_ = ch.QueueBind(q.Name, "", "task_subscribe", false, nil)
	msgs, _ := ch.Consume(q.Name, "", true, false, false, false, nil)

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			fmt.Println("Received a message:", string(d.Body))
		}
	}()

	<-forever
}
