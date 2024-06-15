package common

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

const OrderCreatedEvent = "order.created"

func ConnectAmqp(user, pass, host, port string) (*amqp.Channel, *amqp.Connection) {
	amqpURL := fmt.Sprintf("amqp://%s:%s@%s:%s/", user, pass, host, port)

	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		log.Fatal(err)
	}

	c, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	err = c.ExchangeDeclare("exchange", "direct", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	err = c.ExchangeDeclare(OrderCreatedEvent, "fanout", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	return c, conn
}
