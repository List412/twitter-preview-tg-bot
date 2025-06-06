package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"syscall"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/list412/tweets-tg-bot/internal/clients"
	"github.com/list412/tweets-tg-bot/internal/clients/instagram"
	"github.com/list412/tweets-tg-bot/internal/clients/instagram/graphql"
	"github.com/list412/tweets-tg-bot/internal/clients/instagram/instagramscrapper"
	"github.com/list412/tweets-tg-bot/internal/clients/instagram/instagramscrapper2"
	"github.com/list412/tweets-tg-bot/internal/clients/instagram/profileandmedia"
	"github.com/list412/tweets-tg-bot/internal/clients/instagram/saveinsta1"
	"github.com/list412/tweets-tg-bot/internal/clients/instagram/socialapi1instagram"
	"github.com/list412/tweets-tg-bot/internal/clients/rapidApi"
	tgClient "github.com/list412/tweets-tg-bot/internal/clients/telegram"
	tiktok2 "github.com/list412/tweets-tg-bot/internal/clients/tiktok"
	"github.com/list412/tweets-tg-bot/internal/clients/tiktok/tiktok89"
	"github.com/list412/tweets-tg-bot/internal/clients/tiktok/tiktokscraper7"
	twitter2 "github.com/list412/tweets-tg-bot/internal/clients/twitter"
	"github.com/list412/tweets-tg-bot/internal/clients/twitter/twitterapi45"
	"github.com/list412/tweets-tg-bot/internal/clients/twitter/twttrapi"
	"github.com/list412/tweets-tg-bot/internal/commands"
	"github.com/list412/tweets-tg-bot/internal/config"
	"github.com/list412/tweets-tg-bot/internal/dbConn"
	"github.com/list412/tweets-tg-bot/internal/downloader"
	"github.com/list412/tweets-tg-bot/internal/events/consumer/event-consumer"
	"github.com/list412/tweets-tg-bot/internal/events/telegram"
	logger2 "github.com/list412/tweets-tg-bot/internal/logger"
	metrics2 "github.com/list412/tweets-tg-bot/internal/metrics"
	repository2 "github.com/list412/tweets-tg-bot/internal/storage/share/repository"
	"github.com/list412/tweets-tg-bot/internal/storage/users/repository"
	"github.com/list412/tweets-tg-bot/internal/storage/users/service"
)

func main() {

	logger := slog.New(logger2.ContextHandler{Handler: slog.NewJSONHandler(os.Stdout, nil)})
	slog.SetDefault(logger)

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
	twitterService.RegisterApi(twttrapi.NewService(twttrapiClient, twttrapi.Mapper{Downloader: downloader.Downloader{}}), twitterapi45.NewService(twitterApi45Client))

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

	instagramscrapperClient := instagramscrapper.NewClient(rapidApiClient, cfg.InstagramScrapper.Host)
	instagramscrapperService := instagramscrapper.NewService(instagramscrapperClient)

	//instagrambulkscrapperClient := instagrambulkscrapper.NewClient(rapidApiClient, cfg.InstagramBulkScrapper.Host)
	//instagrambulkscrapperService := instagrambulkscrapper.NewService(instagrambulkscrapperClient)

	saveinstaClient := saveinsta1.NewClient(rapidApiClient, cfg.SaveInsta1.Host)
	saveinstaService := saveinsta1.NewService(saveinstaClient)

	profileandmediaClient := profileandmedia.NewClient(rapidApiClient, cfg.ProfileAndMedia.Host)
	profileandmediaService := profileandmedia.NewService(profileandmediaClient)

	instascrapper2Client := instagramscrapper2.NewClient(rapidApiClient, cfg.InstagramScrapper2.Host)
	instascrapper2Service := instagramscrapper2.NewService(instascrapper2Client)

	grapthqlClient := graphql.NewClient("")
	grapthqlService := graphql.NewService(grapthqlClient)

	instaService := instagram.NewService()
	instaService.RegisterApi(grapthqlService)
	instaService.RegisterApi(instascrapper2Service)
	instaService.RegisterApi(instagramscrapperService)
	instaService.RegisterApi(instagramSocialApiService)
	instaService.RegisterApi(profileandmediaService)
	instaService.RegisterApi(saveinstaService)

	contentProviderManager := clients.NewManager()
	contentProviderManager.RegisterService(commands.TweetCmd, twitterService).
		RegisterService(commands.TikTokCmd, ttService).
		RegisterService(commands.InstagramCmd, instaService)

	eventProcessor := telegram.New(
		tgClient.NewClient(cfg.Telegram.Host, cfg.Telegram.Token),
		contentProviderManager,
		cmdParser,
		usersServ,
		cfg.BotHandler,
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
