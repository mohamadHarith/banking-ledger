package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type RabbitMQ struct {
	User     string `env:"RABBITMQ_DEFAULT_USER"`
	Password string `env:"RABBITMQ_DEFAULT_PASS"`
}

type MongoDB struct {
	User     string `env:"MONGODB_USERNAME"`
	Password string `env:"MONGODB_PASSWORD"`
	Database string `env:"MONGODB_DATABASE"`
}

type configs struct {
	RabbitMQ
	MongoDB
}

var conf configs

func GetConf() configs {
	return conf
}

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	if err := env.Parse(&conf); err != nil {
		panic(err)
	}
}
