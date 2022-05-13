package consumer

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/hasanbakirci/mongodb-kafka-connector/config"
	"github.com/hasanbakirci/mongodb-kafka-connector/internal/mongokafka"
	"log"
	"strings"
)

type KafkaConsumer struct {
	service mongokafka.Service
	client  sarama.Consumer
}

func NewConsume(s mongokafka.Service, c sarama.Consumer) KafkaConsumer {
	return KafkaConsumer{
		service: s,
		client:  c,
	}
}

func (k KafkaConsumer) Consume() {
	partitionConsumer, err := k.client.ConsumePartition(config.Config.TOPIC, 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}
	for msg := range partitionConsumer.Messages() {
		log.Printf("Consumed message offset %d\n", msg.Offset)
		result := strings.Split(strings.Split(string(msg.Value), `"$oid": "`)[1], `"`)[0]
		fmt.Println(result)
		k.service.Create(context.Background(), mongokafka.CreateKafkaLogRequest{LogId: result})
	}
}
