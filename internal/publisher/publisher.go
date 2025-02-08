package publisher

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
	"github.com/google/uuid"
)

type Publisher struct {
	producer sarama.SyncProducer
	topic    string
}

func New(topic string) (*Publisher, error) {
	// Sarama configuration
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true          // Ensure successes are returned
	config.Producer.RequiredAcks = sarama.WaitForAll // Wait for all replicas to acknowledge
	config.Producer.Retry.Max = 5                    // Retry up to 5 times

	// Create a new sync producer
	producer, err := sarama.NewSyncProducer([]string{"kafka:9094"}, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create Kafka producer: %w", err)
	}

	return &Publisher{
		producer: producer,
		topic:    topic,
	}, nil
}

func (p *Publisher) Publish(msg []byte) error {
	// Generate a unique key for the message
	key := uuid.NewString()

	// Create a Kafka message
	kafkaMsg := &sarama.ProducerMessage{
		Topic: p.topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.ByteEncoder(msg),
	}

	// Send the message
	partition, offset, err := p.producer.SendMessage(kafkaMsg)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}

	// Log success
	log.Printf("Message sent successfully! Topic: %s, Partition: %d, Offset: %d\n", p.topic, partition, offset)
	return nil
}

func (p *Publisher) Stop() {
	if err := p.producer.Close(); err != nil {
		log.Printf("error closing Kafka producer: %v", err)
	}
}
