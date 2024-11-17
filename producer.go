package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
)

func ProduceMessage(broker, topic, message string) error {
	writer := kafka.Writer{
		Addr:     kafka.TCP(broker),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}

	defer writer.Close()

	err := writer.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte("key"),
		Value: []byte(message),
	})
	if err != nil {
		return fmt.Errorf("failed to write message: %s", err)
	}
	return nil
}
