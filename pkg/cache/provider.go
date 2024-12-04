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

type SerializationType uint8

const (
	SerializationTypeUnknown = iota
	SerializationTypeJson
	SerializationTypeGob
)

func (c *ObjectCache[T]) createKey(k string) string {
	return c.prefix + k
}

type ObjectCache[T any] struct {
	prefix            string
	provider          Provider
	serializationType SerializationType
}

func NewObjectCache[T any](provider Provider, st SerializationType, prefix string) *ObjectCache[T] {
	return &ObjectCache[T]{
		prefix:            prefix,
		provider:          provider,
		serializationType: st,
	}
}

func NewJsonObjectCache[T any](provider Provider, prefix string) *ObjectCache[T] {
	return &ObjectCache[T]{
		prefix:            prefix,
		provider:          provider,
		serializationType: SerializationTypeJson,
	}
}

func (c *ObjectCache[T]) Set(ctx context.Context, key string, ttl time.Duration, in T) error {
	data, err := c.marshal(in)

	if err != nil {
		return err
	}

	return c.provider.Set(ctx, key, ttl, data)
}

func (c *ObjectCache[T]) Get(ctx context.Context, key string) (T, error) {
	var t = new(T)
	data, err := c.provider.Get(ctx, key)
	if err != nil {
		if errors.Is(err, ErrCacheMiss) {
			return *t, nil
		}
		return *t, err
	}

	return *t, c.unmarshal(data, &t)
}

func (c *ObjectCache[T]) Del(ctx context.Context, key string) error {
	return c.provider.Del(ctx, c.createKey(key))
}
