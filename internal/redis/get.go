package redis

import (
	"context"
	"fmt"
)

func (r *Redis) Get(ctx context.Context, idempotencyKey string) (any, error) {
	key := idempotencyPrefix + idempotencyKey

	val, err := r.redis.Get(ctx, key).Result()
	if err != nil {
		return nil, fmt.Errorf("r.redis.Get: %w", err)
	}

	return val, nil
}
