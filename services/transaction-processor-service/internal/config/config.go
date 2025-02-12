package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type MySql struct {
	User         string `env:"MYSQL_USER"`
	Password     string `env:"MYSQL_PASSWORD"`
	Database     string `env:"MYSQL_DATABASE"`
	RootPassword string `env:"MYSQL_ROOT_PASSWORD"`
}

type configs struct {
	MySql
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
