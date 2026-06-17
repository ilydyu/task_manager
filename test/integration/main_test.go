package integration

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/ilydyu/task_manager.git/config"
	"github.com/ilydyu/task_manager.git/internal/app"
	"github.com/ilydyu/task_manager.git/internal/server"
	"github.com/ilydyu/task_manager.git/pkg/client"
	rep "github.com/ilydyu/task_manager.git/pkg/mysql"
	redisClient "github.com/ilydyu/task_manager.git/pkg/redis"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mysql"
	"github.com/testcontainers/testcontainers-go/modules/redis"
)

var ctx = context.Background()

func Test_Integration(t *testing.T) {
	suite.Run(t, &Suite{})
}

type Suite struct {
	suite.Suite
	*require.Assertions
	containerDB    *mysql.MySQLContainer
	containerCache *redis.RedisContainer
	client         *client.Client
	c              config.Config
}

func (s *Suite) SetupSuite() { // В начале всех тестов
	s.Assertions = s.Require()

	ctx := context.Background()
	mysqlContainer, err := mysql.Run(ctx,
		"mysql:8.0.36",
		mysql.WithDatabase("test"),
		mysql.WithUsername("user"),
		mysql.WithPassword("password"),
	)

	if err != nil {
		panic(err)
	}

	redisContainer, err := redis.Run(ctx,
		"redis:7",
		redis.WithSnapshotting(10, 1),
		redis.WithLogLevel(redis.LogLevelVerbose),
	)

	if err != nil {
		panic(err)
	}

	s.containerDB = mysqlContainer
	s.containerCache = redisContainer

	mysqlHost, _ := mysqlContainer.Host(ctx)
	mysqlPort, _ := mysqlContainer.MappedPort(ctx, "3306/tcp")

	redisHost, _ := redisContainer.Host(ctx)
	redisPort, _ := redisContainer.MappedPort(ctx, "6379/tcp")

	c := config.Config{
		App: config.App{
			Name:    "task_manager",
			Version: "test",
			Secret:  "secret",
		},
		HTTP: server.Config{
			Port: "8080",
		},
		Repository: rep.Config{
			Host:     mysqlHost,
			Port:     mysqlPort.Port(),
			User:     "user",
			Password: "password",
			DBName:   "test",
		},
		Redis: redisClient.Config{
			Addr: fmt.Sprintf("%s:%s", redisHost, redisPort.Port()),
		},
	}

	s.c = c

	s.ResetMigrations()

	log.Logger = zerolog.Nop()

	go func() {
		err := app.Run(ctx, c)
		s.NoError(err)
	}()

	s.client = client.New(client.Config{Host: "localhost", Port: "8080"})

	time.Sleep(time.Second)
}

func (s *Suite) TearDownSuite() {
	err := testcontainers.TerminateContainer(s.containerDB)

	if err != nil {
		log.Printf("failed to terminate mysql container: %s", err)
	}

	err = testcontainers.TerminateContainer(s.containerCache)

	if err != nil {
		log.Printf("failed to terminate redis container: %s", err)
	}
} // В конце всех тестов

func (s *Suite) SetupTest() {} // Перед каждым тестом

func (s *Suite) TearDownTest() {
	s.ResetTables()
} // После каждого теста
