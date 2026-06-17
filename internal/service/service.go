package service

import (
	"context"
	"time"

	"github.com/ilydyu/task_manager.git/internal/domain"
	"github.com/ilydyu/task_manager.git/internal/repository"
)

type Redis interface {
	IsExists(ctx context.Context, idempotencyKey string) bool
	Get(ctx context.Context, idempotencyKey string) (any, error)
	Set(ctx context.Context, idempotencyKey string, value any, ttl time.Duration) error
}

type Mysql interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
	CreateTask(ctx context.Context, task *domain.Task) error
	CreateTeam(ctx context.Context, team *domain.Team) error
	CreateTeamMember(ctx context.Context, member *domain.TeamMember) error
	GetTeamByID(ctx context.Context, id int64) (domain.Team, error)
	GetTeamMember(ctx context.Context, userID int64, teamID int64) (domain.TeamMember, error)
	GetUserByEmail(ctx context.Context, email string) (domain.User, error)
	GetUserByID(ctx context.Context, id int64) (domain.User, error)
	GetUserTeams(ctx context.Context, userID int64) ([]domain.Team, error)
	IsExists(ctx context.Context, table string, id int64) (bool, error)
	IsTeamMemberExists(ctx context.Context, userID int64, teamID int64) (bool, error)
	TrackTaskChanges(ctx context.Context, task domain.Task) error
}

type Service struct {
	repository *repository.Repository
	redis      Redis
	secret     string
}

func New(repository *repository.Repository, redis Redis, secret string) *Service {
	return &Service{
		repository: repository,
		redis:      redis,
		secret:     secret,
	}
}
