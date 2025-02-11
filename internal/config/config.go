package config

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

func NewConfig() (*Config, error) {
	config := Config{}

	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	err = env.Parse(&config)
	if err != nil {
		return nil, err
	}

	return &config, err
}

type Config struct {
	Telegram              Telegram
	Tweeter               Tweeter
	Storage               Storage
	Consumer              Consumer
	Db                    Db
	Admin                 Admin
	Prometheus            Prometheus
	Twttrapi              Twttrapi
	TwitterApi45          TwitterApi45
	TikTokScraper2        TikTokScraper2
	TikTok89              TikTok89
	TikTokScrapper7       TikTokScrapper7
	Socialapi1Instagram   Socialapi1Instagram
	InstagramScrapper     InstagramScrapper
	InstagramBulkScrapper InstagramBulkScrapper
	SaveInsta1            SaveInsta1
	ProfileAndMedia       ProfileAndMedia
	RapidApi              RapidApi
}

type Prometheus struct {
	Port int `env:"PROM_PORT"`
}

type Telegram struct {
	Token string `env:"TELEGRAM_TOKEN"`
	Host  string `env:"TELEGRAM_HOST"`
}

type Tweeter struct {
	Token string `env:"TWITTER_TOKEN"`
	Host  string `env:"TWITTER_HOST"`
}

type Twttrapi struct {
	Host string `env:"TWTTRAPI_HOST"`
}

type TwitterApi45 struct {
	Host string `env:"TWITTERAPI45_HOST"`
}

type TikTokScraper2 struct {
	Host string `env:"TIKTOK_SCRAPPER2_HOST"`
}

type TikTok89 struct {
	Host string `env:"TIKTOK89_HOST"`
}

type TikTokScrapper7 struct {
	Host string `env:"TIKTOK_SCRAPPER7_HOST"`
}

type ProfileAndMedia struct {
	Host string `env:"PROFILEANDMEDIA_HOST"`
}

type Socialapi1Instagram struct {
	Host string `env:"SOCIALAPI1INSTAGRAM_HOST"`
}

type InstagramScrapper struct {
	Host string `env:"INSTAGRAMSCRAPPER_HOST"`
}

type InstagramBulkScrapper struct {
	Host string `env:"INSTAGRAMBULKSCRAPPER_HOST"`
}

type SaveInsta1 struct {
	Host string `env:"SAVEINSTA1_HOST"`
}

type RapidApi struct {
	Token string `env:"RAPIDAPI_KEY"`
}

type Storage struct {
	BasePath string `env:"STORAGE_BASE_PATH"`
}

type Consumer struct {
	BatchSize int `env:"CONSUMER_BATCH_SIZE"`
}

type Db struct {
	User string `env:"DB_USER"`
	Pass string `env:"DB_PASS"`
	Host string `env:"DB_HOST"`
	Port string `env:"DB_PORT"`
	Name string `env:"DB_NAME"`
}

func (d Db) Dsn() string {
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	//	c.user, c.pass, c.host, c.port, c.name)
	dsn := fmt.Sprintf("mongodb://%s:%s@%s:%s", d.User, d.Pass, d.Host, d.Port)
	return dsn
}

type Admin struct {
	Id int `env:"ADMIN_ID"`
}
