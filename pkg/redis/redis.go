package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/huhaophp/hblog/pkg/config"
	"log"
)

var RD *redis.Client

func init() {
	r := config.Conf.Redis
	RD = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", r.Host, r.Port),
		Password: r.Pass,
		DB:       r.DB,
		PoolSize: 10,
	})
	if str, err := RD.Ping(context.Background()).Result(); err != nil || str != "PONG" {
		log.Fatalf("redis connect ping failed, err: %v", err)
	}
}
