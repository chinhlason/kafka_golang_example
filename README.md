# Kafka with GO

## How to run this project

1. Start the Docker container:  
using command : `docker-compose up -d`

2. Create a topic:  
using command : `make create` (you can find the full command in the Makefile)  
you can also check the topic using the command : `make list`

3. Start the producer and consumer:  
using command : `make produce` and `make consume`

4. Run the program:  
using command : `go run main.go`
choose lib you want to test "seg" for segmentio or "sar" for sarama
you can choose "c" for consumer or "p" for producer, check the result in the terminal that you run the command above step 3.

