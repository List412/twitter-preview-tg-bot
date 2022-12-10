package main

import (
	"context"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"syscall"
	tgClient "tweets-tg-bot/internal/clients/telegram"
	"tweets-tg-bot/internal/clients/twitter/twitterScraper"
	"tweets-tg-bot/internal/config"
	"tweets-tg-bot/internal/dbConn"
	"tweets-tg-bot/internal/events/consumer/event-consumer"
	"tweets-tg-bot/internal/events/telegram"
	metrics2 "tweets-tg-bot/internal/metrics"
	repository2 "tweets-tg-bot/internal/storage/share/repository"
	"tweets-tg-bot/internal/storage/users/repository"
	"tweets-tg-bot/internal/storage/users/service"
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

	// Create a new registry.
	reg := prometheus.NewRegistry()

	// Add Go module build info.
	reg.MustRegister(collectors.NewBuildInfoCollector())
	reg.MustRegister(collectors.NewGoCollector(
		collectors.WithGoCollectorRuntimeMetrics(collectors.GoRuntimeMetricsRule{Matcher: regexp.MustCompile("/.*")}),
	))
	metricsHandler := metrics2.Metrics{}
	metricsHandler.Register(reg)

	// Expose the registered metrics via HTTP.
	http.Handle("/metrics", promhttp.HandlerFor(
		reg,
		promhttp.HandlerOpts{
			// Opt into OpenMetrics to support exemplars.
			EnableOpenMetrics: true,
		},
	))

	go func() {
		addr := fmt.Sprintf(":%d", cfg.Prometheus.Port)
		log.Printf("Starting web server at %s\n", addr)
		err := http.ListenAndServe(addr, nil)
		if err != nil {
			log.Printf("http.ListenAndServer: %v\n", err)
		}
	}()

	usersCollection := db.Database(cfg.Db.Name).Collection("users")
	shareCollections := db.Database(cfg.Db.Name).Collection("share")

	err = setupMongo(ctx, db, cfg)
	if err != nil {
		panic(err)
	}

	usersRepo := repository.New(usersCollection)
	shareRepo := repository2.New(shareCollections)
	usersServ := service.New(usersRepo, shareRepo, &metricsHandler, cfg.Admin)

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
