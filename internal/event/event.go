package event

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"github.com/spf13/viper"
	"net"
	"os"
)

const (
	DefaultKafkaHost = "localhost"
	DefaultKafkaPort = "9092"
)

func New(config *viper.Viper, eventName string) (*kafka.Conn, error) {
	host, ok := os.LookupEnv("KAFKA_HOST")
	if !ok {
		host = DefaultKafkaHost
	}

	port, ok := os.LookupEnv("KAFKA_PORT")
	if !ok {
		port = DefaultKafkaPort
	}

	eventName = fmt.Sprintf("%s.%s.%s", "events", eventName, "topic")
	conn, err := kafka.DialLeader(context.Background(), "tcp", net.JoinHostPort(host, port), config.GetString(eventName), 0)
	if err != nil {
		fmt.Printf("Kafka connection error: %s", err.Error())
	}
	defer func(conn *kafka.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Printf("Kafka connection close error: %s", err.Error())
		}
	}(conn)

	return conn, err
}
