package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"strconv"
)

func main() {
	conn, _ := amqp.Dial("amqp://guest:guest@localhost:5672/")
	ch, _ := conn.Channel()
	q, _ := ch.QueueDeclare("task_queue", false, false, false, false, nil)
	for i := 0; i < 10; i++ {
		body := "msg" + strconv.Itoa(i)
		_ = ch.Publish(
			"", q.Name, false, false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		fmt.Println(body)
	}
}
