package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"strconv"
)

func main() {
	conn, _ := amqp.Dial("amqp://guest:guest@localhost:5672/")
	ch, _ := conn.Channel()
	q1, _ := ch.QueueDeclare("q1", false, false, false, false, nil)
	q2, _ := ch.QueueDeclare("q2", false, false, false, false, nil)
	for i := 0; i < 5; i++ {
		body := "msg" + strconv.Itoa(i)
		_ = ch.Publish(
			"", q1.Name, false, false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		fmt.Println("Sent message: " + body)
	}

	msgs, _ := ch.Consume(q2.Name, "", true, false, false, false, nil)

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Println("Received a message:", string(d.Body))
		}
	}()
	<-forever

}
