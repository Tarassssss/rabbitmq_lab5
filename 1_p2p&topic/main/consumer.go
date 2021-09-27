package main

import (
	"fmt"
	"time"
	"github.com/streadway/amqp"
	"github.com/nothingelsematters7/golang_proto/utils"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	utils.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	ch, err := conn.Channel()
	utils.FailOnError(err, "Failed to open a channel")
	defer ch.Close()
	q, err := ch.QueueDeclare("task_queue", false, false, false, false, nil)
	utils.FailOnError(err, "Failed to declare a queue")
	msgs, _ := ch.Consume(q.Name, "", true, false, false, false, nil)

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Println("Received a message: ", string(d.Body))
			time.Sleep(100*time.Millisecond)
		}
	}()
	<-forever
}
