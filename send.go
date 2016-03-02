package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatal("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func main() {
	rabbitURL := os.Getenv("RABBIT_URL")
	conn, err := amqp.Dial(rabbitURL)
	failOnError(err, "failed to connect to rabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // autoDelete
		false,   // exclusive
		false,   // noWait
		nil,     // args
	)

	failOnError(err, "Failed to declare a queue")

	m := Message{"Alice", "Hello"}
	b, err := json.Marshal(m)

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        b,
		})
	log.Printf(" [x] Sent %s", b)
	failOnError(err, "Failed to publish a message")
}
