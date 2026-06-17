package redis

import (
	"context"
	"fmt"
	"time"
)

func (r *Redis) Set(ctx context.Context, idempotencyKey string, value any, ttl time.Duration) error {
	key := idempotencyPrefix + idempotencyKey

	err := r.redis.Set(ctx, key, value, ttl).Err()

	if err != nil {
		return fmt.Errorf("redis: Set: r.redis.Set: %w", err)
	}

	return nil
}
