package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"strconv"
)

func main() {
	conn, _ := amqp.Dial("amqp://guest:guest@localhost:5672/")
	ch, _ := conn.Channel()
	q, _ := ch.QueueDeclare("durableQueue", true, false, false, false, nil)
	for i := 0; i < 10; i++ {
		body := "test message" + strconv.Itoa(i)
		err := ch.Publish(
			"", q.Name, false, false,
			amqp.Publishing{
				ContentType:  "text/plain",
				Body:         []byte(body),
				DeliveryMode: 2,
			})
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(body)
		}
	}
}
