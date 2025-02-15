package repository

import (
	"context"
	"strconv"
)

func (r *Repository) SetUserBalance(ctx context.Context, userId, accountId string, balance uint32) error {
	return r.redis.HSet(ctx, userId, accountId, strconv.Itoa(int(balance))).Err()
}

func (r *Repository) GetUserBalance(ctx context.Context, userId, accountId string) (uint32, error) {
	balance, err := r.redis.HGet(ctx, userId, accountId).Uint64()
	if err != nil {
		return 0, err
	}

	return uint32(balance), nil
}
