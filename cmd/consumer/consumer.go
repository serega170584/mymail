package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":     "localhost:9092",
		"broker.address.family": "v4",
		"group.id":              "my-topic-group",
		"session.timeout.ms":    6000,
		"auto.offset.reset":     "earliest",
	})

	if err != nil {
		fmt.Printf("Failed to create consumer: %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("Created Consumer %v\n", c)

	topics := []string{"my-topic"}
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
				fmt.Printf("%% Message on %s:\n%s\n",
					e.TopicPartition, string(e.Value))
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
