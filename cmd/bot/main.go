package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"os/signal"
	"syscall"
	tgClient "tweets-tg-bot/internal/clients/telegram"
	"tweets-tg-bot/internal/clients/twitter/twitterScraper"
	"tweets-tg-bot/internal/config"
	"tweets-tg-bot/internal/consumer/event-consumer"
	"tweets-tg-bot/internal/dbConn"
	"tweets-tg-bot/internal/events/telegram"
	repository2 "tweets-tg-bot/internal/share/repository"
	"tweets-tg-bot/internal/users/repository"
	"tweets-tg-bot/internal/users/service"
)

func main() {

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	ctx, cancel := signal.NotifyContext(ctx, os.Kill, os.Interrupt, syscall.SIGTERM)
	defer cancel()

	db, err := dbConn.Open(ctx, cfg.Db)
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close(ctx, db)

	scrapper := twitterScraper.NewScrapper()

	usersCollection := db.Database(cfg.Db.Name).Collection("users")
	shareCollections := db.Database(cfg.Db.Name).Collection("share")

	err = setupMongo(ctx, db, cfg)
	if err != nil {
		panic(err)
	}

	usersRepo := repository.New(usersCollection)
	shareRepo := repository2.New(shareCollections)
	usersServ := service.New(usersRepo, shareRepo, cfg.Admin)

	eventProcessor := telegram.New(
		tgClient.NewClient(cfg.Telegram.Host, cfg.Telegram.Token),
		scrapper,
		usersServ,
	)

	go scrapper.UpdateTokenJob()
	go eventProcessor.HandleUsers()

	consumer := event_consumer.NewConsumer(eventProcessor, eventProcessor, cfg.Consumer.BatchSize)

	log.Printf("service started")

	if err := consumer.Start(ctx); err != nil {
		log.Fatal(err)
	}
}

func setupMongo(ctx context.Context, db *mongo.Client, cfg *config.Config) error {
	uniqueUsers := true
	usersCollection := db.Database(cfg.Db.Name).Collection("users")
	shareCollections := db.Database(cfg.Db.Name).Collection("share")

	indexModel := mongo.IndexModel{
		Keys: bson.M{
			"userName": 1, // index in ascending order
		}, Options: &options.IndexOptions{Unique: &uniqueUsers},
	}
	_, err := usersCollection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		return err
	}
	_, err = shareCollections.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		return err
	}
	return nil
}
