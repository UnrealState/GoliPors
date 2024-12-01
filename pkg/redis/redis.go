package redis

import (
	"fmt"
	"golipors/config"

	"github.com/go-redis/redis"
)

var Client *redis.Client

func InitRedis(cfg config.RedisConfig) error {
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	Client = redis.NewClient(&redis.Options{
		Addr: addr,
	})
	_, err := Client.Ping().Result()
	return err
}
