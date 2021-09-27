package main

import (
	"fmt"
	"github.com/streadway/amqp"
)

func main() {
	conn, _ := amqp.Dial("amqp://guest:guest@localhost:5672/")
	ch, _ := conn.Channel()
	q1, _ := ch.QueueDeclare("q1", false, false, false, false, nil)
	q2, _ := ch.QueueDeclare("q2", false, false, false, false, nil)
	msgs, _ := ch.Consume(q1.Name, "", true, false, false, false, nil)

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			text := string(d.Body) + " (edited)"
			_ = ch.Publish(
				"", q2.Name, false, false,
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        []byte(text),
				})
			fmt.Println("Received and resent a message:", string(d.Body))
		}
	}()
	<-forever
}
