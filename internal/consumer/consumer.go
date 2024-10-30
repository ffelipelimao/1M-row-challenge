package consumer

import (
	"context"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Consumer struct {
	ctx      context.Context
	consumer *kafka.Consumer
	topic    string
	handler  func(ctx context.Context, msg []byte) error
}

func New(ctx context.Context, topic string, handler func(ctx context.Context, msg []byte) error) (*Consumer, error) {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		"client.id":         "consumer-xx",
		"group.id":          "group-xx",
	}

	consumer, err := kafka.NewConsumer(configMap)
	if err != nil {
		return nil, err
	}

	topics := []string{topic}
	consumer.SubscribeTopics(topics, nil)

	return &Consumer{
		ctx:      ctx,
		consumer: consumer,
		topic:    topic,
		handler:  handler,
	}, nil
}

func (c *Consumer) Start() {
	for {

		//-1 The call will block for at most `timeout` waiting for
		// a new message or error. `timeout` may be set to -1 for
		// indefinite wait.

		msg, err := c.consumer.ReadMessage(-1)
		if err != nil {
			fmt.Println("error to process message", err)
		}
		err = c.handler(c.ctx, msg.Value)
		if err != nil {
			fmt.Println("error to handle a message", err)
		}
	}
}

func (c *Consumer) Stop() {
	c.consumer.Close()
}
