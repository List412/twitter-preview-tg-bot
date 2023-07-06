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
	Telegram   Telegram
	Tweeter    Tweeter
	Storage    Storage
	Consumer   Consumer
	Db         Db
	Admin      Admin
	Prometheus Prometheus
	Twttrapi   Twttrapi
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
	Host  string `env:"TWTTRAPI_HOST"`
	Token string `env:"TWTTRAPI_KEY"`
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
