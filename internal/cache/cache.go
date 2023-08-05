package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	redis2 "github.com/hhandhuan/ku-bbs/pkg/redis"
	"time"
)

func Cache() *cache {
	return &cache{store: redis2.GetInstance()}
}

type cache struct {
	store *redis.Client
}

func (c *cache) Get(key string, obj interface{}) error {
	return c.store.Get(context.Background(), key).Scan(&obj)
}

func (c *cache) Set(key string, obj interface{}, ttl time.Duration) error {
	return c.store.Set(context.Background(), key, obj, ttl).Err()
}

func (c *cache) Del(key string) error {
	return c.store.Del(context.Background(), key).Err()
}
