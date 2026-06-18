package dto

import "github.com/ilydyu/task_manager.git/internal/domain"

type GetTeamStatsOutput struct {
	Data []domain.TeamStatistics `json:"data"`
}
