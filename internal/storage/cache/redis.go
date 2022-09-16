package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/rs/zerolog/log"

	"github.com/kostaspt/go-tiny/config"
)

var DefaultExpiration = 1 * time.Hour

type Redis struct {
	Client *redis.Client
}

func NewRedis(c *config.Config) (*Redis, func()) {
	cl := redis.NewClient(&redis.Options{
		Addr:     c.Cache.Addr,
		Username: c.Cache.Username,
		Password: c.Cache.Password,
	})

	cleanup := func() {
		if err := cl.Close(); err != nil {
			log.Err(err).Send()
			return
		}
	}

	return &Redis{Client: cl}, cleanup
}

func (c Redis) Get(ctx context.Context, key string, data any) error {
	d, err := c.Client.Get(ctx, key).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil
		}
		return err
	}

	return json.Unmarshal(d, &data)
}

func (c Redis) Set(ctx context.Context, key string, data any, ttl ...time.Duration) error {
	d, err := json.Marshal(data)
	if err != nil {
		return err
	}

	expiredTime := DefaultExpiration
	if len(ttl) > 0 {
		expiredTime = ttl[0]
	}

	return c.Client.Set(ctx, key, d, expiredTime).Err()
}

func (c Redis) Delete(ctx context.Context, key string) error {
	return c.Client.Del(ctx, key).Err()

}

func (c Redis) Close() error {
	return c.Client.Close()
}
