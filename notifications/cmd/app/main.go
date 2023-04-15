package main

import (
	"context"
	"log"
	"route256/libs/logger"
	"route256/notifications/config"
	kafka_client "route256/notifications/internal/clients/kafka"
	"route256/notifications/internal/domain"
	"time"

	"go.uber.org/zap"
)

var (
	develMode = false
)

func main() {
	err := config.Init()
	if err != nil {
		log.Fatal("config init", err)
	}

	// logger
	log := *logger.New(develMode)

	handle := func(key string, value string, time time.Time) {
		log.Info("Message claimed", zap.String("key", key), zap.String("value", value), zap.Time("time", time))
	}

	logsReciver := kafka_client.NewReciver(config.ConfigData.KafkaTopic, config.ConfigData.KafkaBrokers, handle)

	service := domain.New(log, logsReciver)

	service.Subscribe(context.Background())
}
