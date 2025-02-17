package config

import (
	"os"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type RabbitMQ struct {
	ServiceName string `env:"RABBITMQ_SERVICE_NAME"`
	User        string `env:"RABBITMQ_DEFAULT_USER"`
	Password    string `env:"RABBITMQ_DEFAULT_PASS"`
}

type MongoDB struct {
	ServiceName string `env:"MONGODB_SERVICE_NAME"`
	User        string `env:"MONGODB_USERNAME"`
	Password    string `env:"MONGODB_PASSWORD"`
	Database    string `env:"MONGODB_DATABASE"`
}

type configs struct {
	RabbitMQ
	MongoDB
	ServiceName string `env:"TRANSACTION_LOGGER_SERVICE_NAME"`
	ServicePort string `env:"TRANSACTION_LOGGER_SERVICE_PORT"`
	Env         string `env:"ENV" envDefault:"local"`
}

func (conf configs) IsLocalEnvironment() bool {
	return conf.Env == "local"
}

func (conf configs) IsDevelopmentEnvironment() bool {
	return conf.Env == "dev"
}

var conf configs

func GetConf() configs {
	return conf
}

func init() {
	if os.Getenv("ENV") != "dev" {
		err := godotenv.Load("../../.env")
		if err != nil {
			panic(err)
		}
	}

	if err := env.Parse(&conf); err != nil {
		panic(err)
	}
}
