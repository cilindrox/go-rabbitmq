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

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // autoAck
		false,  // exclusive
		false,  // noLocal
		false,  // noWait
		nil,    // args
	)

	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			var m Message
			if err := json.Unmarshal(d.Body, &m); err != nil {
				log.Printf("Error parsing msg", err)
			}
			SendSMS(m)
		}
	}()

	log.Println(" [*] Waiting for messages. Press CTRL + C to exit")
	<-forever
}
