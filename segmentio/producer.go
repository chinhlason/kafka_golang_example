package segmentio

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"time"
)

func ProduceMessage(broker, topic, message string) error {
	start := time.Now()

	writer := kafka.Writer{
		Addr:         kafka.TCP(broker),
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		BatchTimeout: 10 * time.Millisecond,
		BatchSize:    100,
	}

	defer func() {
		fmt.Printf("time taken to produce message: %v\n", time.Since(start))
		writer.Close()
	}()

	err := writer.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte("key"),
		Value: []byte(message),
	})
	if err != nil {
		return fmt.Errorf("failed to write message: %s", err)
	}
	return nil
}
