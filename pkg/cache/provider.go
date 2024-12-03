package cache

import (
	"context"
	"errors"
	"time"
)

var (
	ErrCacheMiss = errors.New("cache miss")
)

type Provider interface {
	Set(ctx context.Context, key string, ttl time.Duration, data []byte) error
	Get(ctx context.Context, key string) ([]byte, error)
	Del(ctx context.Context, key string) error
}
