package main

import (
	"common"
	"encoding/json"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	amqpUser     = "host"
	amqpPassword = "host"
	amqpHost     = "localhost"
	amqpPort     = "5672"
)

func main() {
	c, conn := common.ConnectAmqp(amqpUser, amqpPassword, amqpHost, amqpPort)
	defer func() {
		conn.Close()
		c.Close()
	}()

	q, err := c.QueueDeclare(common.OrderCreatedEvent, true, false, false, false, nil)

	if err != nil {
		log.Fatal(err)
	}

	msgch, err := c.Consume(q.Name, "", true, false, false, false, nil)

	donech := make(chan struct{})

	go listen(msgch, donech)

	fmt.Println("Listening to AMQP Queue: ", q.Name)

	<-donech
}

func listen(msgch <-chan amqp.Delivery, donech chan<- struct{}) {
	defer func() {
		donech <- struct{}{}
	}()
	for msg := range msgch {
		log.Printf("Received a message: %s", msg.Body)
		order := &common.Order{}
		if err := json.Unmarshal(msg.Body, order); err != nil {
			// It is possible to configure a dead-letter queue or even requeue the message if second parameter is true
			msg.Nack(false, false)
			log.Printf("Failed to unmarshal message body: %v", err)
			continue
		}
		log.Printf("Payment done for message: %s", msg.Body)
	}
}
