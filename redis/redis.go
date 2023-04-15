package redis

import (
	"fmt"

	"github.com/Wave-ETH-Global/wave-node/config"
	"github.com/go-redis/redis/v8"
)

func ProvideRedis(cfg *config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf(
			"%s:%d",
			cfg.Redis.Host,
			cfg.Redis.Port,
		),
		DB:       int(cfg.Redis.Database),
		Password: "",
	})
}
