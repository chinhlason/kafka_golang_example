package segmentio

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
)

func ConsumeMessages(broker, topic, groupId string) error {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker},
		Topic:   topic,
		GroupID: groupId,
	})
	defer reader.Close()

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			return fmt.Errorf("failed to read message: %s", err)
		}
		fmt.Printf("message received: %s\n", string(msg.Value))
	}
}
