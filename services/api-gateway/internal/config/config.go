package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type RabbitMQ struct {
	ServiceName string `env:"RABBITMQ_SERVICE_NAME"`
	User        string `env:"RABBITMQ_DEFAULT_USER"`
	Password    string `env:"RABBITMQ_DEFAULT_PASS"`
}

type Redis struct {
	ServiceName string `env:"REDIS_SERVICE_NAME"`
	Password    string `env:"REDIS_PASSWORD"`
}

type AuthenticationService struct {
	ServiceName string `env:"AUTHENTICATION_SERVICE_NAME"`
	ServicePort string `env:"AUTHENTICATION_SERVICE_PORT"`
}

type TransactionLoggerService struct {
	ServiceName string `env:"TRANSACTION_LOGGER_SERVICE_NAME"`
	ServicePort string `env:"TRANSACTION_LOGGER_SERVICE_PORT"`
}

type TransactionProcessorService struct {
	ServiceName string `env:"TRANSACTION_PROCESSOR_SERVICE_NAME"`
	ServicePort string `env:"TRANSACTION_PROCESSOR_SERVICE_PORT"`
}

type configs struct {
	RabbitMQ
	Redis
	AuthenticationService
	TransactionLoggerService
	TransactionProcessorService
	ServiceName string `env:"API_GATEWAY_SERVICE_NAME"`
	ServicePort string `env:"API_GATEWAY_PORT"`
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
	err := godotenv.Load("../../.env")
	if err != nil {
		panic(err)
	}

	if err := env.Parse(&conf); err != nil {
		panic(err)
	}
}
