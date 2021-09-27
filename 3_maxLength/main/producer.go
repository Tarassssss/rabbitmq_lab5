package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"strconv"
	"github.com/nothingelsematters7/golang_proto/utils"
)

func main() {
	conn, _ := amqp.Dial("amqp://guest:guest@localhost:5672/")
	ch, err := conn.Channel()
	utils.FailOnError(err, "Failed to open a channel")
	defer ch.Close()
	q, _ := ch.QueueDeclare("limitedQueue5", false, false, false, false, amqp.Table{"x-max-length": 10, "x-overflow": "reject-publish"})
	q.Messages = 10
	for i := 0; i < 20; i++ {
		body := "msg" + strconv.Itoa(i)
		err := ch.Publish(
			"", q.Name, false, false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(body)
		}
	}
}
