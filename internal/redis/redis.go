package redis

import (
	"time"

	"github.com/ilydyu/task_manager.git/pkg/redis"
)

const (
	idempotencyPrefix = "task_manager:idempotency:"
	ttl               = 5 * time.Minute
)

type Redis struct {
	redis *redis.Client
}

func New(client *redis.Client) *Redis {
	return &Redis{
		redis: client,
	}
}
