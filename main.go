package main

import (
	"fmt"
	"kafka/sarama"
	"kafka/segmentio"
	"os"
)

func main() {
	broker := "localhost:29092"
	topic := "test"

	fmt.Println("enter lib: segmentio or sarama")
	var lib string
	fmt.Scanln(&lib)

	fmt.Println("enter mode: producer or consumer")
	var mode string
	fmt.Scanln(&mode)

	switch mode {
	case "p":
		fmt.Println("enter message")
		var message string
		fmt.Scanln(&message)
		if lib == "seg" {
			err := segmentio.ProduceMessage(broker, topic, message)
			if err != nil {
				fmt.Println("error producing message: ", err)
				os.Exit(1)
			}
		} else {
			err := sarama.ProduceMessage(broker, topic, message)
			if err != nil {
				fmt.Println("error producing message: ", err)
				os.Exit(1)
			}
		}
	case "c":
		groupId := "test-group"
		if lib == "seg" {
			err := segmentio.ConsumeMessages(broker, topic, groupId)
			if err != nil {
				fmt.Println("error consuming message: ", err)
				os.Exit(1)
			}
		} else {
			err := sarama.ConsumeMessage(broker, topic, groupId)
			if err != nil {
				fmt.Println("error consuming message: ", err)
				os.Exit(1)
			}
		}
	default:
		fmt.Println("invalid mode")
		os.Exit(1)
	}
}
