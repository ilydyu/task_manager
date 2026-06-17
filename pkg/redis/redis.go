package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

type Config struct {
	Addr     string `envconfig:"REDIS_ADDR" required:"true"`
	Password string `envconfig:"REDIS_PASSWORD"`
	DB       int    `envconfig:"REDIS_DB" default:"0"`
}

type Client struct {
	*redis.Client
}

func New(ctx context.Context, c Config) (*Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     c.Addr,
		Password: c.Password,
		DB:       c.DB,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("client.Ping: %w", err)
	}

	return &Client{Client: client}, nil
}

func (c *Client) Close() {
	err := c.Client.Close()
	if err != nil {
		log.Error().Err(err).Msg("redis: close")
	}

	log.Info().Msg("redis: closed")
}
