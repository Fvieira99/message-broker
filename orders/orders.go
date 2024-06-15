package main

import (
	"common"
	"context"
	"encoding/json"
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

	order, err := json.Marshal(common.Order{
		Id: "order_1",
		Items: []common.Item{
			{
				Id:       "item_1",
				Quantity: 1,
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	err = c.PublishWithContext(context.Background(), "", q.Name, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        order,
	})

	if err != nil {
		log.Fatal(err)
	}
	log.Println("Order Published")
}
