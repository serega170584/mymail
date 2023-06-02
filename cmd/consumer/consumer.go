package main

import (
	"awesomeProject/internal/config"
	notifier "awesomeProject/internal/proto"
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"gopkg.in/gomail.v2"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	const (
		defaultKafkaHost = "localhost"
		defaultKafkaPort = "9092"
	)

	if len(os.Args) != 4 {
		fmt.Printf("Usage: %s <broker> <group> <timeoutms..>\n", os.Args[0])
		os.Exit(1)
	}

	conf := config.New()

	kafkaHost, ok := os.LookupEnv("KAFKA_HOST")
	if !ok {
		kafkaHost = defaultKafkaHost
	}

	kafkaPort, ok := os.LookupEnv("KAFKA_PORT")
	if !ok {
		kafkaPort = defaultKafkaPort
	}

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":     net.JoinHostPort(kafkaHost, kafkaPort),
		"broker.address.family": "v4",
		"group.id":              os.Args[1],
		"session.timeout.ms":    os.Args[3],
		"auto.offset.reset":     "earliest",
	})

	if err != nil {
		fmt.Printf("Failed to create consumer: %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("Created Consumer %v\n", c)

	topics := []string{os.Args[2]}
	err = c.SubscribeTopics(topics, nil)
	if err != nil {
		fmt.Printf("Subscribe topic event attempt error: %s\n", err.Error())
		os.Exit(1)
	}

	run := true

	for run {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev := c.Poll(100)
			if ev == nil {
				continue
			}

			switch e := ev.(type) {
			case *kafka.Message:
				fmt.Printf("%% Message on %s:\n%s\n", e.TopicPartition, string(e.Value))

				var in *notifier.EmailRequest
				err = json.Unmarshal(e.Value, &in)
				if err != nil {
					fmt.Printf("Unmararshal email request: %s", err.Error())
					run = false
					break
				}

				message := gomail.NewMessage()
				message.SetHeader("From", conf.GetString("mail.from"))
				message.SetHeader("To", in.To...)
				message.SetHeader("Subject", fmt.Sprintf("grpc handler was triggered at %s", time.Now().String()))
				// TODO: google mailchimp если сложно то найдем другое решение
				// dialer := gomail.NewDialer(server.config.GetString("mail.host"), server.config.GetInt("mail.port"), server.config.GetString("mail.from"), "111")
				// dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
				// if err := dialer.DialAndSend(message); err != nil {
				// 	log.Printf("failed to send mail: %s\n", err.Error())
				// 	return &emptypb.Empty{}, err
				// }
				log.Println("email is sent")
				if e.Headers != nil {
					fmt.Printf("%% Headers: %v\n", e.Headers)
				}
			case kafka.Error:
				fmt.Printf("%% Error: %v: %v\n", e.Code())
				if e.Code() == kafka.ErrAllBrokersDown {
					run = false
				}
			default:
				fmt.Printf("Ignored %v\n", e)
			}
		}
	}

	fmt.Printf("Closing consumer\n")
	err = c.Close()
	if err != nil {
		fmt.Printf("Closing consumer error\n")
	}
}
