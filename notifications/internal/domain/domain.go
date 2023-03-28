package domain

import "context"

type KafkaReciver interface {
	Subscribe(context.Context) error
}

type Service struct {
	LogsReciver KafkaReciver
}

func New(logsReciver KafkaReciver) *Service {
	return &Service{
		logsReciver,
	}
}
