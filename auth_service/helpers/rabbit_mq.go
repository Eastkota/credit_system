package helpers

import(
	"auth_service/config"
	"fmt"

	"github.com/rabbitmq/amqp091-go"
)

func connectRabbitMQ() (*amqp091.Connection, error) {
	conn, err := amqp091.Dial(config.RabbitMQURL())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %v", err)
	}
	return conn, nil
}

func createRabbitMQChannel() (*amqp091.Channel, error) {
	conn, err := connectRabbitMQ()
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		defer conn.Close()
		return nil, fmt.Errorf("failed to open a channel: %v", err)
	}
	return ch, nil
}

func RabbitMQPublishEvent(name string, exchangeType string, body []byte) error {
	ch, err := createRabbitMQChannel()
	if err != nil {
		return err
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		name,   // name
		exchangeType,        // type
		true,           // durable
		false,          // auto-deleted
		false,          // internal
		false,          // no-wait
		nil,           // arguments
	)
	if err != nil {
		return err
	}
	err = ch.Publish(
		name,
		name,
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		return err
	}
	return nil
}

func RabbitMQConsumeEvent(name string, exchangeType string) (<-chan amqp091.Delivery, *amqp091.Channel, error) {
	ch, err := createRabbitMQChannel()
	if err != nil {
		return nil, nil, err
	}

	err = ch.ExchangeDeclare(
		name,   // name
		exchangeType,        // type
		true,           // durable
		false,          // auto-deleted
		false,          // internal
		false,          // no-wait
		nil,           // arguments
	)
	if err != nil {
		return nil, nil,err
	}

	queue, err := ch.QueueDeclare(
		name,  // name
		true,   // durable
		false,  // delete when unused
		false,  // exclusive
		false,  // no-wait
		nil,    // arguments
	)
	if err != nil {
		ch.Close()
		return nil, nil, err
	}

	// The queue is bound to the exchange to receive messages
	err = ch.QueueBind(
		queue.Name, // queue name
		name,       // routing key
		name,       // exchange
		false,      // no-wait
		nil,        // arguments
	)
	if err != nil {
		ch.Close()
		return nil, nil, err
	}

	// Start consuming messages
	msgs, err := ch.Consume(
		queue.Name, // queue
		"",         // consumer
		false,      // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	if err != nil {
		ch.Close()
		return nil, nil, err
	}

	return msgs, ch, nil
}