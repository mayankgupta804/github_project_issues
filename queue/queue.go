package queue

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

// This package level variable will hold the connection to our RabbitMQ instance
var conn *amqp.Connection

// Init initialize the package level "conn" variable that represents the connection the the Rabbitmq server
func Init(c string) {
	var err error
	conn, err = amqp.Dial(c)
	if err != nil {
		log.Fatalf("Could not connect to Rabbitmq server: %v", err)
		panic(err)
	}
}

// Publish publishes the message to the RabbitMQ queue for consumption by the worker
func Publish(q string, msg []byte) error {
	// Create a channel through which we publish
	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()
	// Create the payload with the message that we specify in the arguments
	payload := amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "application/json",
		Body:         msg,
	}
	// Publish the message to the queue specified in the arguments
	if err := ch.Publish("", q, false, false, payload); err != nil {
		return fmt.Errorf("[Publish] failed to publish to queue %v", err)
	}
	return nil
}

// Subscribe is used by the worker to subscribe to messages published to the RabbitMQ queue for consumption
// and subsequent processing
func Subscribe(qName string) (<-chan amqp.Delivery, func(), error) {
	// Create a channel through which we publish
	ch, err := conn.Channel()
	if err != nil {
		return nil, nil, err
	}
	// assert that the queue exists (creates a queue if it doesn't)
	q, err := ch.QueueDeclare(qName, false, false, false, false, nil)
	// create a channel in go, through which incoming messages will be received
	c, err := ch.Consume(q.Name, "", false, false, false, false, nil)
	// return the created channel
	return c, func() { ch.Close() }, err
}
