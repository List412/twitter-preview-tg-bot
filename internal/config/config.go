package config

import (
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
	Telegram Telegram
	Tweeter  Tweeter
	Storage  Storage
	Consumer Consumer
}

type Telegram struct {
	Token string `env:"TELEGRAM_TOKEN"`
	Host  string `env:"TELEGRAM_HOST"`
}

type Tweeter struct {
	Token string `env:"TWITTER_TOKEN"`
	Host  string `env:"TWITTER_HOST"`
}

type Storage struct {
	BasePath string `env:"STORAGE_BASE_PATH"`
}

type Consumer struct {
	BatchSize int `env:"CONSUMER_BATCH_SIZE"`
}
