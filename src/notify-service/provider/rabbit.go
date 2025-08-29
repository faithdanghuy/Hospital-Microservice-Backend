package provider

import (
	"context"
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

type RabbitSubscriber struct {
	url     string
	timeout time.Duration
	conn    *amqp.Connection
	ch      *amqp.Channel
}

func NewRabbitSubscriber(url string, timeout time.Duration) *RabbitSubscriber {
	return &RabbitSubscriber{url: url, timeout: timeout}
}

// StartConsume connects and consumes messages from a durable "notify" queue.
// handler should process and return nil on success.
func (r *RabbitSubscriber) StartConsume(handler func(queue string, body []byte) error) error {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	r.conn, err = amqp.Dial(r.url)
	if err != nil {
		return err
	}
	r.ch, err = r.conn.Channel()
	if err != nil {
		_ = r.conn.Close()
		return err
	}

	q, err := r.ch.QueueDeclare("notify", true, false, false, false, nil)
	if err != nil {
		return err
	}

	msgs, err := r.ch.Consume(q.Name, "", false, false, false, false, nil)
	if err != nil {
		return err
	}

	// provider/rabbit_subscriber.go
	for d := range msgs {
		if err := handler(q.Name, d.Body); err != nil {
			fmt.Printf("notify: handler error: %v\n", err)
			// If it’s already been redelivered, drop it (no requeue) to avoid loops.
			if d.Redelivered {
				_ = d.Nack(false, false)
			} else {
				_ = d.Nack(false, true) // requeue once
			}
			continue
		}
		_ = d.Ack(false)
	}

	// context deadline or closed channel handling optional
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		return nil
	}
}
