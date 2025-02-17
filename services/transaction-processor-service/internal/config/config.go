package config

import (
	"github.com/caarlos0/env/v11"
)

type MySql struct {
	ServiceName  string `env:"MYSQL_SERVICE_NAME"`
	User         string `env:"MYSQL_USER"`
	Password     string `env:"MYSQL_PASSWORD"`
	Database     string `env:"MYSQL_DATABASE"`
	RootPassword string `env:"MYSQL_ROOT_PASSWORD"`
}

type RabbitMQ struct {
	ServiceName string `env:"RABBITMQ_SERVICE_NAME"`
	User        string `env:"RABBITMQ_DEFAULT_USER"`
	Password    string `env:"RABBITMQ_DEFAULT_PASS"`
}

type configs struct {
	MySql
	RabbitMQ
	ServiceName string `env:"TRANSACTION_PROCESSOR_SERVICE_NAME"`
	ServicePort string `env:"TRANSACTION_PROCESSOR_SERVICE_PORT"`
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
