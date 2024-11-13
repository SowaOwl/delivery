package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type Cache interface {
	Get(key string) (string, error)
	Set(key string, value string) error
}

type RedisCache struct {
	client *redis.Client
}

func newRedisCache(client *redis.Client) *RedisCache {
	return &RedisCache{client: client}
}

func (r *RedisCache) Get(key string) (string, error) {
	return r.client.Get(context.Background(), key).Result()
}

func (r *RedisCache) Set(key string, value string) error {
	return r.client.Set(context.Background(), key, value, 10*10*60).Err()
}
