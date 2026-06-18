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
	CreateTask(ctx context.Context, task *domain.Task) error
	CreateTeam(ctx context.Context, team *domain.Team) error
	CreateTeamMember(ctx context.Context, member *domain.TeamMember) error
	CreateUser(ctx context.Context, user *domain.User) error
	GetInvalidAssignments(ctx context.Context) ([]domain.InvalidAssignment, error)
	GetTaskByID(ctx context.Context, id int64) (domain.Task, error)
	GetTaskHistory(ctx context.Context, taskID int64) ([]domain.TaskHistory, error)
	GetTasks(ctx context.Context, teamID string, assigneeID string, status string, cursor string) ([]domain.Task, error)
	GetTeamMember(ctx context.Context, userID int64, teamID int64) (domain.TeamMember, error)
	GetTeamStats(ctx context.Context) ([]domain.TeamStatistics, error)
	GetTopCreators(ctx context.Context) ([]domain.TopCreatorStats, error)
	GetUserByEmail(ctx context.Context, email string) (domain.User, error)
	GetUserTeams(ctx context.Context, userID int) ([]domain.Team, error)
	IsTeamMemberExists(ctx context.Context, userID int64, teamID int64) (bool, error)
	TrackTaskChanges(ctx context.Context, task domain.Task, action string, oldValue []byte) error
	UpdateTask(ctx context.Context, task *domain.Task) error
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
