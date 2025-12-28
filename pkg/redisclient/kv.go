package redisclient

import (
	"context"
	"time"
)

func (c *Client) Get(ctx context.Context, key string) (string, error) {
	return c.rdb.Get(ctx, key).Result()
}

func (c *Client) Set(ctx context.Context, key, value string, ttl time.Duration) error {
	return c.rdb.Set(ctx, key, value, ttl).Err()
}

func (c *Client) Delete(ctx context.Context, key string) error {
	return c.rdb.Del(ctx, key).Err()
}

func (c *Client) Incr(ctx context.Context, key string, ttl time.Duration) (int64, error) {
	val, err := c.rdb.Incr(ctx, key).Result()
	if err != nil {
		return 0, err
	}

	if val == 1 && ttl > 0 {
		_ = c.rdb.Expire(ctx, key, ttl).Err()
	}

	return val, nil
}
