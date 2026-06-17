package cache

import (
	"BackendPOS/internal/config"
	"context"

	"github.com/redis/go-redis/v9"
)

var (
	Rdb *redis.Client
	Ctx = context.Background()
)

func InitRedis(cfg *config.Config) *redis.Client {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     cfg.RedisHost + ":" + cfg.RedisPort,
		Password: cfg.RedisPassword,
		DB:       0,
		Username: cfg.RedisUsername,
	})

	if err := Rdb.Ping(Ctx).Err(); err != nil {

		panic("failed to connect redis: " + err.Error())
	}
	return Rdb

}
