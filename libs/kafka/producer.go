package kafka

import (
	"github.com/Shopify/sarama"
)

type Producer interface {
	SendMessage(*sarama.ProducerMessage) (partition int32, offset int64, err error)
}

func NewAsyncProducer(brokers []string) (sarama.AsyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewHashPartitioner // sarama.NewRoundRobinPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Idempotent = true // exactly once
	config.Net.MaxOpenRequests = 1

	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true

	producer, err := sarama.NewAsyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}

	// либо тут либо в обертке делаем, но обязательно обрабатывать ошикбки/ сообщения
	// // config.Producer.Return.Errors = true
	// go func() {
	// 	for e := range producer.Errors() {
	// 		fmt.Println(e.Msg.Key, e.Error())
	// 	}
	// }()

	// // config.Producer.Return.Successes = true
	// go func() {
	// 	for m := range producer.Successes() {
	// 		fmt.Println(m.Key)
	// 	}
	// }()

	return producer, nil
}
