package publisher

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/google/uuid"
)

type Publisher struct {
	publisher           *kafka.Producer
	producerEventChanel chan kafka.Event
	topic               string
}

func New(topic string) (*Publisher, error) {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers":   "kafka:9092",
		"delivery.timeout.ms": "0",
		"acks":                "all",
	}

	publisherKafka, err := kafka.NewProducer(configMap)
	if err != nil {
		fmt.Println("Error to create producer:", err)
		return nil, err
	}

	return &Publisher{
		publisher:           publisherKafka,
		producerEventChanel: make(chan kafka.Event),
		topic:               topic,
	}, nil
}

func (p *Publisher) Publish(msg []byte) error {
	key := []byte(uuid.NewString())

	kafkaMsg := &kafka.Message{
		Key:   key,
		Value: msg,
		TopicPartition: kafka.TopicPartition{
			Topic:     &p.topic,
			Partition: kafka.PartitionAny,
		},
	}

	err := p.publisher.Produce(kafkaMsg, p.producerEventChanel)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	go handleEventMessage(p.producerEventChanel)

	return nil
}

func (p *Publisher) Stop() {
	p.publisher.Close()
}

func handleEventMessage(channel chan kafka.Event) {
	for e := range channel {
		switch ev := e.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				fmt.Println("error to send message")
			}
		}
	}
}
