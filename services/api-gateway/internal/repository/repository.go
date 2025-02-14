package repository

import (
	"context"

	"github.com/mohamadHarith/banking-ledger/services/api-gateway/internal/config"
	"github.com/redis/go-redis/v9"
)

type Repository struct {
	redis *redis.Client
}

func New() *Repository {
	conf := config.GetConf()

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: conf.Redis.Password,
		DB:       0,
	})

	if client == nil {
		panic("nil redis client")
	}

	return &Repository{
		redis: client,
	}
}

func (r *Repository) SetUserBalance(ctx context.Context, userId string, balance uint32) error {
	return r.redis.Set(ctx, userId, balance, 0).Err()
}

// func (r *Repository) GetUserBalance(ctx context.Context, userId string, balance uint32) error {
// 	return r.redis.Set(ctx, userId, balance, 0).Err()
// }
