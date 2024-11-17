package main

import (
	"fmt"
	"os"
)

func main() {
	broker := "localhost:29092"
	topic := "test"

	fmt.Println("enter mode: producer or consumer")
	var mode string
	fmt.Scanln(&mode)

	switch mode {
	case "p":
		fmt.Println("enter message")
		var message string
		fmt.Scanln(&message)
		err := ProduceMessage(broker, topic, message)
		if err != nil {
			fmt.Println("error producing message: ", err)
			os.Exit(1)
		}
	case "c":
		groupId := "test-group"
		err := ConsumeMessages(broker, topic, groupId)
		if err != nil {
			fmt.Println("error consuming messages: ", err)
			os.Exit(1)
		}
	default:
		fmt.Println("invalid mode")
		os.Exit(1)
	}
}
