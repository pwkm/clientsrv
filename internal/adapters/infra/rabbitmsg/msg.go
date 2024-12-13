package rabbitmsg

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/pwkm/clientsrv/internal/core/domain"
	"github.com/pwkm/clientsrv/internal/utils/env"
	"github.com/streadway/amqp"
)

type Message struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

// ------------------------------------
// Kafka stream
// ------------------------------------
type RabbitStream struct {
	Con     *amqp.Connection
	Channel *amqp.Channel
	Q       *amqp.Queue
}

func NewRabbitStream(env *env.Env) (*RabbitStream, error) {
	// Connect to RabbitMQ server
	conn, err := amqp.Dial(env.RabbitCon)
	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ: ", err)
	}
	// defer conn.Close()

	// Create a channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Failed to open a channel: ", err)
	}
	// defer ch.Close()

	// Declare a queue to send to
	queue, err := ch.QueueDeclare(
		env.RabbitQueue, // name of the queue
		false,           // durable
		false,           // delete when unused
		false,           // exclusive
		false,           // no-wait
		nil,             // arguments
	)
	if err != nil {
		log.Fatal("Failed to declare a queue: ", err)
	}
	return &RabbitStream{
		Con:     conn,
		Channel: ch,
		Q:       &queue,
	}, nil
}

func (p *RabbitStream) SendMessage(client *domain.Client) error {

	message, _ := json.Marshal(client)

	err := p.Channel.Publish(
		"",       // exchange
		p.Q.Name, // routing key
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		log.Fatal("Sent message error: %w", err)
	}
	fmt.Println("Message send")

	return err
}
