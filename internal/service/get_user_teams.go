package service

import (
	"context"
	"fmt"

	"github.com/ilydyu/task_manager.git/internal/dto"
)

func (s *Service) GetUserTeams(ctx context.Context, id int) (dto.GetUserTeamsOutput, error) {
	var output dto.GetUserTeamsOutput

	teams, err := s.repository.GetUserTeams(ctx, id)

	if err != nil {
		return output, fmt.Errorf("s.repository.GetUserTeams: %w", err)
	}

	return dto.GetUserTeamsOutput{Teams: teams}, nil
}
