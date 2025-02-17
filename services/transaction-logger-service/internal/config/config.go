package config

import (
	"github.com/caarlos0/env/v11"
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
}

var conf configs

func GetConf() configs {
	return conf
}

func init() {
	// err := godotenv.Load()
	// if err != nil {
	// 	panic(err)
	// }

	if err := env.Parse(&conf); err != nil {
		panic(err)
	}
}
