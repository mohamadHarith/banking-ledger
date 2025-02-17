package repository

import (
	"context"
	"fmt"
	"sync"

	"github.com/mohamadHarith/banking-ledger/services/api-gateway/internal/config"
	"github.com/redis/go-redis/v9"
)

type Repository struct {
	redis *redis.Client
}

var repo *Repository
var once sync.Once

func New() *Repository {
	once.Do(func() {
		conf := config.GetConf()

		redisHost := conf.Redis.ServiceName
		if conf.IsLocalEnvironment() {
			redisHost = "localhost"
		}

		client := redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%v:6379", redisHost),
			Password: conf.Redis.Password,
			DB:       0,
		})

		if client == nil {
			panic("nil redis client")
		}

		err := client.Ping(context.Background()).Err()
		if err != nil {
			panic(err)
		}

		repo = &Repository{
			redis: client,
		}
	})

	return repo
}

func (r *Repository) Close() {
	r.Close()
}
