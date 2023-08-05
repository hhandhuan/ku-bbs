package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/hhandhuan/ku-bbs/pkg/config"
	"github.com/hhandhuan/ku-bbs/pkg/logger"
)

var instance *redis.Client

func GetInstance() *redis.Client {
	return instance
}
func Initialize(conf *config.Redis) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Host, gconv.Int(conf.Port)),
		Password: conf.Pass,
		DB:       gconv.Int(conf.DB),
		PoolSize: 10,
	})
	if str, err := client.Ping(context.Background()).Result(); err != nil || str != "PONG" {
		logger.GetInstance().Fatal().Msgf("redis connect ping failed, err: %v", err)
	}
	instance = client
}
