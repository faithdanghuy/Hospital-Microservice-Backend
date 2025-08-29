package provider

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/rabbitmq/amqp091-go"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitPublisher struct {
	url     string
	timeout time.Duration
	conn    *amqp.Connection
	ch      *amqp.Channel
}

func NewRabbitPublisher(url string, timeout time.Duration) (*RabbitPublisher, error) {
	var conn *amqp.Connection
	var ch *amqp.Channel
	var err error

	deadline := time.Now().Add(timeout)

	for time.Now().Before(deadline) {
		conn, err = amqp.Dial(url)
		if err == nil {
			ch, err = conn.Channel()
			if err == nil {
				_, err = ch.QueueDeclare("notify", true, false, false, false, nil)
				if err == nil {
					return &RabbitPublisher{url: url, timeout: timeout, conn: conn, ch: ch}, nil
				}
			}
		}

		log.Printf("⚠️ RabbitMQ not ready: %v. Retrying in 5s...", err)
		time.Sleep(5 * time.Second)
	}

	return nil, fmt.Errorf("could not connect to RabbitMQ after %s: %w", timeout, err)
}

func (p *RabbitPublisher) Publish(ctx context.Context, body []byte) error {
	return p.ch.PublishWithContext(ctx,
		"",
		"notify",
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
}

func (p *RabbitPublisher) Close() {
	_ = p.ch.Close()
	_ = p.conn.Close()
}
