package domain

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/Shopify/sarama"
)

type logsSender struct {
	asyncProducer sarama.AsyncProducer
	topic         string
}

type Handler func(id string)

func NewOrderLogSender(producer sarama.AsyncProducer, topic string, onSuccess, onFailed Handler) *logsSender {
	s := &logsSender{
		asyncProducer: producer,
		topic:         topic,
	}

	// config.Producer.Return.Errors = true
	go func() {
		for e := range producer.Errors() {
			bytes, _ := e.Msg.Key.Encode()

			onFailed(string(bytes))
			fmt.Println(e.Msg.Key, e.Error())
		}
	}()

	// config.Producer.Return.Successes = true
	go func() {
		for m := range producer.Successes() {
			bytes, _ := m.Key.Encode()

			onSuccess(string(bytes))
			log.Printf("order id: %s, partition: %d, offset: %d", string(bytes), m.Partition, m.Offset)
		}
	}()

	return s
}

func (s *logsSender) SendOrderStatusAsync(order Order) {
	// variant 1 (json)
	bytes, err := json.Marshal(order)
	if err != nil {
		return
	}

	msg := &sarama.ProducerMessage{
		Topic:     s.topic,
		Partition: -1,
		Value:     sarama.ByteEncoder(bytes),
		Key:       sarama.StringEncoder(rune(order.OrderId)),
		Timestamp: time.Now(),
		Headers: []sarama.RecordHeader{
			{
				Key:   []byte("test-header-key"),
				Value: []byte("test-header-value"),
			},
		},
	}

	s.asyncProducer.Input() <- msg
}
