package domain

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/Shopify/sarama"
	"go.uber.org/zap"
)

type logsSender struct {
	asyncProducer sarama.AsyncProducer
	topic         string
}

type Handler func(id string)

func NewOrderLogSender(log *zap.Logger, producer sarama.AsyncProducer, topic string, onSuccess, onFailed Handler) *logsSender {
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
			log.Info("log sended", zap.String("value", string(bytes)), zap.Int32("partition", m.Partition), zap.Int64("offset", m.Offset))
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
