package cache

import (
	"context"
	"time"

	"github.com/nick6969/gopkg/gob"
)

func SetGOBModel[T any](ctx context.Context, key string, model *T, expire time.Duration, s *Cache) error {
	value := gob.Container[T]{RawValue: *model}
	return s.client.SetModel(ctx, key, value, expire)
}

func GetGOBModel[T any](ctx context.Context, key string, s *Cache) (*T, error) {
	var value gob.Container[T]
	err := s.client.GetModel(ctx, key, &value)
	if err != nil {
		return nil, err
	}
	return &value.RawValue, nil
}
