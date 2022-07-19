package service

import (
	"context"
	v8Redis "github.com/go-redis/redis/v8"
	"github.com/huhaophp/hblog/pkg/redis"
	"time"
)

func Cache() *cache {
	return &cache{storage: redis.RD}
}

type cache struct {
	storage *v8Redis.Client
}

func (s *cache) Get(key string) (string, error) {
	if res, err := s.storage.Get(context.Background(), key).Result(); err != nil && err != v8Redis.Nil {
		return "", err
	} else {
		return res, nil
	}
}

func (s *cache) Has(key string) (bool, error) {
	str, err := s.Get(key)
	if err != nil {
		return false, err
	}
	if str == "" || str == "0" {
		return false, nil
	}
	if str == "1" {
		return true, nil
	}

	return false, nil
}

func (s *cache) Set(key, value string, duration time.Duration) error {
	return s.storage.SetEX(context.Background(), key, value, duration).Err()
}

func (s *cache) Forever(key, value string) error {
	return s.storage.Set(context.Background(), key, value, 0).Err()
}

func (s *cache) Forget(key string) error {
	return s.storage.Del(context.Background(), key).Err()
}
