package event

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"net"
	"os"
)

const (
	DefaultKafkaHost = "localhost"
	DefaultKafkaPort = "9092"
)

type Event struct {
	producer *kafka.Producer
}

func New() (*Event, error) {
	host, ok := os.LookupEnv("KAFKA_HOST")
	if !ok {
		host = DefaultKafkaHost
	}

	port, ok := os.LookupEnv("KAFKA_PORT")
	if !ok {
		port = DefaultKafkaPort
	}

	broker := net.JoinHostPort(host, port)

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": broker})

	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err.Error())
		return nil, err
	}

	return &Event{producer: p}, nil
}

func (event *Event) Produce(topic string, message string) error {
	deliveryChan := make(chan kafka.Event)

	p := event.producer
	err := p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}, deliveryChan)

	if err != nil {
		fmt.Printf("Producer deliver channel error: %s", err.Error())
		return err
	}

	e := <-deliveryChan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
	} else {
		fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
			*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	}

	close(deliveryChan)

	return nil
}
