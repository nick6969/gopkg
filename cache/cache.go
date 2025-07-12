package cache

import (
	"context"
	"time"
)

type RedisClient interface {
	Set(ctx context.Context, key string, value string, expiration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Del(ctx context.Context, key string) error
	GetModel(ctx context.Context, key string, model interface{}) error
	SetModel(ctx context.Context, key string, model interface{}, expiration time.Duration) error
	Incr(ctx context.Context, key string) (int64, error)
	Expire(ctx context.Context, key string, expiration time.Duration) error
}

type Cache struct {
	client RedisClient
}

func NewService(client RedisClient) *Cache {
	return &Cache{
		client: client,
	}
}

func (c *Cache) Set(ctx context.Context, key string, value string, expiration time.Duration) error {
	return c.client.Set(ctx, key, value, expiration)
}

func (c *Cache) Get(ctx context.Context, key string) (string, error) {
	return c.client.Get(ctx, key)
}

func (c *Cache) Del(ctx context.Context, key string) error {
	return c.client.Del(ctx, key)
}

func (c *Cache) Incr(ctx context.Context, key string) (int64, error) {
	return c.client.Incr(ctx, key)
}
func (c *Cache) Expire(ctx context.Context, key string, expiration time.Duration) error {
	return c.client.Expire(ctx, key, expiration)
}
