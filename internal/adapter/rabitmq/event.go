package rabbitmq

import (
	ampq "github.com/rabbitmq/amqp091-go"
)

func declareExchange(ch *ampq.Channel) error {
	return ch.ExchangeDeclare(
		"logs_topic", // name
		"topic",      // type
		true,         // durable
		false,        // auto-delete
		false,        // interfal
		false,        // no-wait
		nil,          // arguments
	)
}

func declareRandomQueue(ch *ampq.Channel) (ampq.Queue, error) {
	return ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
}
