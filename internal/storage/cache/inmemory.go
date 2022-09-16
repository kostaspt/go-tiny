package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/patrickmn/go-cache"
)

type InMemory struct {
	cache *cache.Cache
}

func NewInMemory() *InMemory {
	return &InMemory{
		cache: cache.New(5*time.Minute, 10*time.Minute),
	}
}

func (c *InMemory) Get(ctx context.Context, key string, data any) error {
	d, found := c.cache.Get(key)
	if !found {
		return nil
	}

	return json.Unmarshal(d.([]byte), &data)
}

func (c *InMemory) Set(ctx context.Context, key string, data any, ttl ...time.Duration) error {
	d, err := json.Marshal(data)
	if err != nil {
		return err
	}

	expiredTime := DefaultExpiration
	if len(ttl) > 0 {
		expiredTime = ttl[0]
	}

	c.cache.Set(key, d, expiredTime)

	return nil
}

func (c *InMemory) Delete(ctx context.Context, key string) error {
	c.cache.Delete(key)
	return nil
}

func (c *InMemory) Close() error {
	c.cache.Flush()
	return nil
}
