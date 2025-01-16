package initialize

import (
	"context"
	"fmt"

	"github.com/ducnv194023/shoppe_be_go/global"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ctx = context.Background()

func InitRedis() {
	r := global.Config.RedisSetting
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", r.Host, r.Port),
        Password: r.Password,
		DB:       r.DB,
		PoolSize: r.PoolSize,
	})

	_, err := rdb.Ping(ctx).Result()

	if err != nil {
		global.Logger.Error("Failed to connect to Redis", zap.Error(err))
		panic(err)
	}

	global.Logger.Info("Connected to Redis")
	global.Rdb = rdb

	RedisExample()
}

func RedisExample() {
	err := global.Rdb.Set(ctx, "score", 100, 0).Err()

	if err != nil {
		global.Logger.Error("Failed to set score", zap.Error(err))
		panic(err)
	}

	val, err := global.Rdb.Get(ctx, "score").Result()

	if err != nil {
		global.Logger.Error("Failed to get score", zap.Error(err))
		panic(err)
	}

	global.Logger.Info("Score", zap.String("score", val))
}
