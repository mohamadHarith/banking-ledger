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

type configs struct {
	ServiceName string `env:"AUTHENTICATION_SERVICE_NAME"`
	ServicePort string `env:"AUTHENTICATION_SERVICE_PORT"`
	MySql
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
