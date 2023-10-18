package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"golang.org/x/net/context"
)

type Emitter struct {
	connection *amqp.Connection
}

func (e *Emitter) setup() error {
	channel, err := e.connection.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()
	return declareExchange(channel)
}

// Push event to event broker with custom severity
func (e *Emitter) Push(ctx context.Context, event string, severity string) error {
	channel, err := e.connection.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	err = channel.PublishWithContext(
		ctx,
		"logs_topic",
		severity,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(event),
		},
	)
	if err != nil {
		return err
	}
	return nil
}

// NewEmitter create rabbitmq event emitter and setup channel connection
func NewEmitter(connection *amqp.Connection) (Emitter, error) {
	emitter := Emitter{
		connection: connection,
	}

	err := emitter.setup()
	if err != nil {
		return Emitter{}, err
	}

	return emitter, nil
}
