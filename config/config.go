package config

import (
	"fmt"

	"github.com/ilydyu/task_manager.git/internal/logger"
	"github.com/ilydyu/task_manager.git/internal/server"
	"github.com/ilydyu/task_manager.git/pkg/mysql"
	"github.com/ilydyu/task_manager.git/pkg/redis"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type App struct {
	Name    string `envconfig:"APP_NAME" required:"true"`
	Version string `envconfig:"APP_VERSION" required:"true"`
	Env     string `envconfig:"ENV" required:"true"`
	Secret  string `envconfig:"SECRET" required:"true"`
}

type Config struct {
	App        App
	HTTP       server.Config
	Repository mysql.Config
	Logger     logger.Config
	Redis      redis.Config
}

func New() (Config, error) {
	var config Config

	err := godotenv.Load(".env")

	if err != nil {
		return config, fmt.Errorf("godotenv.Load: %w", err)
	}

	err = envconfig.Process("", &config)

	if err != nil {
		return config, fmt.Errorf("envconfig.Process: %w", err)
	}

	return config, nil
}
