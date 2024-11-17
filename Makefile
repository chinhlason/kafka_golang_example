TOPIC = test
HOST = kafka_test

create:
	docker exec -it $(HOST) kafka-topics --create --bootstrap-server localhost:29092 --replication-factor 1 --partitions 1 --topic $(TOPIC)

list:
	docker exec $(HOST) kafka-topics --describe --topic $(TOPIC) --bootstrap-server localhost:29092

produce:
	 docker exec -it $(HOST) kafka-console-producer --bootstrap-server localhost:29092 --topic $(TOPIC) --request-required-acks 1

consume:
	docker exec $(HOST) kafka-console-consumer --bootstrap-server localhost:29092 --topic $(TOPIC) --from-beginning