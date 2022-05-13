package cmd

import (
	"fmt"
	"github.com/hasanbakirci/mongodb-kafka-connector/internal/consumer"
	"github.com/hasanbakirci/mongodb-kafka-connector/internal/mongokafka"
	"github.com/hasanbakirci/mongodb-kafka-connector/pkg/kafkaclient"
	"github.com/hasanbakirci/mongodb-kafka-connector/pkg/mongoclient"
)

type listener struct {
	consume consumer.KafkaConsumer
}

func NewListener() listener {
	db, err := mongoclient.ConnectDb()
	if err != nil {
		fmt.Println("Db connection error")
	}

	kafka := kafkaclient.NewKafkaClient("localhost:9092")
	repository := mongokafka.NewKafkalogRepository(db)
	service := mongokafka.NewService(repository)
	consume := consumer.NewConsume(service, kafka)

	return listener{consume: consume}
}

func (l listener) StartListener() {
	go l.consume.Consume()
}
