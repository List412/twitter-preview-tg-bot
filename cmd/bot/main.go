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
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"syscall"
	"tweets-tg-bot/internal/clients/instagram"
	"tweets-tg-bot/internal/clients/instagram/socialapi1instagram"
	"tweets-tg-bot/internal/clients/rapidApi"
	tgClient "tweets-tg-bot/internal/clients/telegram"
	tiktok2 "tweets-tg-bot/internal/clients/tiktok"
	"tweets-tg-bot/internal/clients/tiktok/tiktok89"
	"tweets-tg-bot/internal/clients/tiktok/tiktokscraper7"
	twitter2 "tweets-tg-bot/internal/clients/twitter"
	"tweets-tg-bot/internal/clients/twitter/twitterapi45"
	"tweets-tg-bot/internal/clients/twitter/twttrapi"
	"tweets-tg-bot/internal/commands"
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
		slog.Error("error parsing config", "error", err)
	}

	ctx := context.Background()
	ctx, cancel := signal.NotifyContext(ctx, os.Kill, os.Interrupt, syscall.SIGTERM)
	defer cancel()

	db, err := dbConn.Open(ctx, cfg.Db)
	if err != nil {
		slog.Error("error opening database connection", "error", err)
	}
	defer dbConn.Close(ctx, db)

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
		slog.Info("starting web server", "port", addr)
		err := http.ListenAndServe(addr, nil)
		if err != nil {
			slog.Error("http.ListenAndServe", "error", err)
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

	rapidApiClient := rapidApi.NewClient(cfg.RapidApi.Token)

	twttrapiClient := twttrapi.NewClient(rapidApiClient, cfg.Twttrapi.Host)
	twitterApi45Client := twitterapi45.NewClient(rapidApiClient, cfg.TwitterApi45.Host)

	twitterService := twitter2.NewService()
	twitterService.RegisterApi(twttrapi.NewService(twttrapiClient), twitterapi45.NewService(twitterApi45Client))

	twitterCmdParser := twitter2.CommandParser{}
	tiktokCmdParser := tiktok2.CommandParser{}
	instaCmdParser := instagram.CommandParser{}

	cmdParser := commands.Parsers{}
	cmdParser.RegisterParser(commands.TweetCmd, twitterCmdParser)
	cmdParser.RegisterParser(commands.TikTokCmd, tiktokCmdParser)
	cmdParser.RegisterParser(commands.InstagramCmd, instaCmdParser)

	tiktokClient := tiktok89.NewClient(rapidApiClient, cfg.TikTok89.Host)
	tiktokService := tiktok89.NewService(tiktokClient)

	tiktok7Client := tiktokscraper7.NewClient(rapidApiClient, cfg.TikTokScrapper7.Host)
	tiktok7Service := tiktokscraper7.NewService(tiktok7Client)

	ttService := tiktok2.NewService()
	ttService.RegisterApi(tiktokService, tiktok7Service)

	instagramSocialApiClient := socialapi1instagram.NewClient(rapidApiClient, cfg.Socialapi1Instagram.Host)
	instagramSocialApiService := socialapi1instagram.NewService(instagramSocialApiClient)

	instaService := instagram.NewService()
	instaService.RegisterApi(instagramSocialApiService)

	eventProcessor := telegram.New(
		tgClient.NewClient(cfg.Telegram.Host, cfg.Telegram.Token),
		twitterService,
		ttService,
		instaService,
		cmdParser,
		usersServ,
	)

	go eventProcessor.HandleUsers()

	consumer := event_consumer.NewConsumer(eventProcessor, eventProcessor, cfg.Consumer.BatchSize)

	slog.Info("starting consumer")

	if err := consumer.Start(ctx); err != nil {
		slog.Error("consumer.Start", "error", err)
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
