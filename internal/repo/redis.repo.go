package repo

import (
	"encoding/json"
	"context"
	"time"
	"github.com/redis/go-redis/v9"
)
type IRedisRepo interface {
    SetKey(
		ctx context.Context, 
		key string, 
		value interface{}, 
		expiry time.Duration,
	) (*redis.StatusCmd, error)
}

type RedisRepo struct {
    client *redis.Client
}

func (rr *RedisRepo) SetKey(ctx context.Context, key string, value interface{}, expiry time.Duration) (*redis.StatusCmd, error) {
    jsonValue, err := json.Marshal(value)
    if err != nil {
        return nil, err
    }
    return rr.client.Set(ctx, key, jsonValue, expiry), nil
}

func NewRedisRepo(client *redis.Client) IRedisRepo {
	return &RedisRepo{
		client: client,
	}
}