package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type RabbitMQ struct {
	User     string `env:"RABBITMQ_DEFAULT_USER"`
	Password string `env:"RABBITMQ_DEFAULT_PASS"`
}

type Redis struct {
	Password string `env:"REDIS_PASSWORD"`
}

type configs struct {
	RabbitMQ
	Redis
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
