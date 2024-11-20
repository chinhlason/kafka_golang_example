package sarama

import (
	"fmt"
	"github.com/IBM/sarama"
	"time"
)

func ProduceMessage(broker, topic, message string) error {
	start := time.Now()
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.Retry.Max = 5

	producer, err := sarama.NewSyncProducer([]string{broker}, config)
	if err != nil {
		return err
	}
	defer func() {
		fmt.Printf("time taken to produce message: %v\n", time.Since(start))
		_ = producer.Close()
	}()

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		return err
	}

	fmt.Println("message sent to partition", partition, "with offset", offset)

	return nil
}
