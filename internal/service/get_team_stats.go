package service

import (
	"context"
	"fmt"

	"github.com/ilydyu/task_manager.git/internal/dto"
)

func (s *Service) GetTeamStats(ctx context.Context) (dto.GetTeamStatsOutput, error) {
	var output dto.GetTeamStatsOutput

	stats, err := s.repository.GetTeamStats(ctx)

	if err != nil {
		return output, fmt.Errorf("s.repository.GetStatsTeams: %w", err)
	}

	return dto.GetTeamStatsOutput{Data: stats}, nil
}
