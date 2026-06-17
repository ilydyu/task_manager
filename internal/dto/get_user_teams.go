package dto

import (
	"github.com/ilydyu/task_manager.git/internal/domain"
)

type GetUserTeamsOutput struct {
	Teams []domain.Team
}
