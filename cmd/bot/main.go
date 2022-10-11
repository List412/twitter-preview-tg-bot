package main

import (
	"log"
	tgClient "tweets-tg-bot/internal/clients/telegram"
	twimg_cdn "tweets-tg-bot/internal/clients/twitter/twimg-cdn"
	"tweets-tg-bot/internal/config"
	"tweets-tg-bot/internal/consumer/event-consumer"
	"tweets-tg-bot/internal/events/telegram"
)

func main() {

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	eventProcessor := telegram.New(
		tgClient.NewClient(cfg.Telegram.Host, cfg.Telegram.Token),
		twimg_cdn.NewClient(cfg.Tweeter.Host),
	)

	consumer := event_consumer.NewConsumer(eventProcessor, eventProcessor, cfg.Consumer.BatchSize)

	log.Printf("service started")

	if err := consumer.Start(); err != nil {
		log.Fatal(err)
	}
}
