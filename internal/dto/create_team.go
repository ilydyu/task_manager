package dto

import (
	"time"

	"github.com/ilydyu/task_manager.git/internal/domain"
)

type CreateTeamInput struct {
	UserID int64  `json:"user_id"`
	Name   string `json:"name"`
}

type CreateTeamOutput struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Members   []domain.TeamMember
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
