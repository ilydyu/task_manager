package service

import (
	"context"
	"fmt"

	"github.com/ilydyu/task_manager.git/internal/domain"
	"github.com/ilydyu/task_manager.git/internal/dto"
)

func (s *Service) CreateTeam(ctx context.Context, input dto.CreateTeamInput) (dto.CreateTeamOutput, error) {
	var output dto.CreateTeamOutput

	team, err := domain.NewTeam(input.UserID, input.Name)

	if err != nil {
		return output, fmt.Errorf("domain.NewTeam: %w", err)
	}

	err = s.repository.CreateTeam(ctx, &team)

	if err != nil {
		return output, fmt.Errorf("s.repository.CreateTeam: %w", err)
	}

	member, err := domain.NewTeamMember(input.UserID, team.ID, domain.RoleOwner)

	if err != nil {
		return output, fmt.Errorf("domain.NewTeamMember: %w", err)
	}

	err = s.repository.CreateTeamMember(ctx, &member)

	if err != nil {
		return output, fmt.Errorf("s.repository.CreateTeamMember: %w", err)
	}

	return dto.CreateTeamOutput{
		ID:        team.ID,
		Name:      team.Name,
		CreatedAt: team.CreatedAt,
		Members:   []domain.TeamMember{member},
	}, nil
}
