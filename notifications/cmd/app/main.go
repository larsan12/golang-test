package main

import (
	"context"
	"log"
	"route256/notifications/config"
	kafka_client "route256/notifications/internal/clients/kafka"
	"route256/notifications/internal/domain"
	"time"
)

func main() {
	err := config.Init()
	if err != nil {
		log.Fatal("config init", err)
	}

	handle := func(key string, value string, time time.Time) {
		log.Printf("Message claimed: key = %s, value = %s, timestamp = %v", key, value, time)
	}

	logsReciver := kafka_client.NewReciver(config.ConfigData.KafkaTopic, config.ConfigData.KafkaBrokers, handle)

	service := domain.New(logsReciver)

	service.Subscribe(context.Background())
}
