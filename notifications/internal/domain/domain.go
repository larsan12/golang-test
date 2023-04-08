package domain

import (
	"context"

	"go.uber.org/zap"
)

type KafkaReciver interface {
	Subscribe(context.Context) error
}

type Service struct {
	log         zap.Logger
	LogsReciver KafkaReciver
}

func New(log zap.Logger, logsReciver KafkaReciver) *Service {
	return &Service{
		log,
		logsReciver,
	}
}
