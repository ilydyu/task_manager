package main

import (
	"context"

	"github.com/ilydyu/task_manager.git/config"
	"github.com/ilydyu/task_manager.git/internal/app"
	"github.com/ilydyu/task_manager.git/internal/logger"
	"github.com/rs/zerolog/log"
)

func main() {
	c, err := config.New()

	if err != nil {
		log.Fatal().Err(err).Msg("config.New")
	}

	logger.Init(c.Logger)

	ctx := context.Background()

	err = app.Run(ctx, c)

	if err != nil {
		log.Error().Err(err).Msg("app.Run")
	}
}
