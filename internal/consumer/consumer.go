package consumer

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

type Consumer struct {
	ctx      context.Context
	consumer sarama.Consumer
	topic    string
	handler  func(ctx context.Context, msg []byte) error
}

func New(ctx context.Context, topic string, handler func(ctx context.Context, msg []byte) error) (*Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	consumer, err := sarama.NewConsumer([]string{"kafka:9092"}, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create Kafka consumer: %w", err)
	}

	return &Consumer{
		ctx:      ctx,
		consumer: consumer,
		topic:    topic,
		handler:  handler,
	}, nil
}

func (c *Consumer) Start() {
	partitions, err := c.consumer.Partitions(c.topic)
	if err != nil {
		log.Fatalf("failed to get partitions for topic %s: %v", c.topic, err)
	}

	for _, partition := range partitions {
		pc, err := c.consumer.ConsumePartition(c.topic, partition, sarama.OffsetNewest)
		if err != nil {
			log.Fatalf("failed to consume partition %d: %v", partition, err)
		}

		go func(pc sarama.PartitionConsumer) {
			defer pc.AsyncClose()

			for {
				select {
				case msg := <-pc.Messages():
					// Handle the message
					err := c.handler(c.ctx, msg.Value)
					if err != nil {
						log.Printf("error handling message: %v", err)
					}
				case err := <-pc.Errors():
					// Log any errors
					log.Printf("error consuming message: %v", err)
				case <-c.ctx.Done():
					// Stop consuming if the context is canceled
					return
				}
			}
		}(pc)
	}
}

func (c *Consumer) Stop() {
	if err := c.consumer.Close(); err != nil {
		log.Printf("error closing Kafka consumer: %v", err)
	}
}
