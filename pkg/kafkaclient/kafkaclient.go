package kafkaclient

import (
	sarama "github.com/Shopify/sarama"
)

func NewKafkaClient(url string) sarama.Consumer {
	consumer, err := sarama.NewConsumer([]string{url}, sarama.NewConfig())
	if err != nil {
		panic(err)
	}
	return consumer
}
