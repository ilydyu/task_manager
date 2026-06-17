package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/ilydyu/task_manager.git/config"
	"github.com/ilydyu/task_manager.git/internal/controller/http"
	"github.com/ilydyu/task_manager.git/internal/redis"
	"github.com/ilydyu/task_manager.git/internal/repository"
	"github.com/ilydyu/task_manager.git/internal/server"
	"github.com/ilydyu/task_manager.git/internal/service"
	"github.com/ilydyu/task_manager.git/pkg/mysql"
	redislib "github.com/ilydyu/task_manager.git/pkg/redis"
	"github.com/ilydyu/task_manager.git/pkg/router"
	"github.com/ilydyu/task_manager.git/pkg/transaction"
	"github.com/rs/zerolog/log"
)

func Run(ctx context.Context, c config.Config) error {
	pool, err := mysql.New(ctx, c.Repository)

	if err != nil {
		return fmt.Errorf("mysql.New: %w", err)
	}

	transaction.Init(pool)

	redisClient, err := redislib.New(ctx, c.Redis)

	if err != nil {
		return fmt.Errorf("redislib.New: %w", err)
	}

	s := service.New(repository.New(), redis.New(redisClient), c.App.Secret)

	r := router.New()
	http.Router(r, s, c.App.Secret)
	httpServer := server.New(r, c.HTTP)

	log.Info().Msg("App started")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	<-sig

	log.Info().Msg("App got signal to stop")

	httpServer.Close()

	pool.Close()

	redisClient.Close()

	log.Info().Msg("App stopped")

	return nil
}
