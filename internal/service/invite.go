package service

import (
	"context"
	"fmt"

	"github.com/ilydyu/task_manager.git/internal/domain"
	"github.com/ilydyu/task_manager.git/internal/dto"
)

func (s *Service) Invite(ctx context.Context, input dto.InviteInput) (dto.InviteOutput, error) {
	var output dto.InviteOutput
	invitedMember, err := s.repository.GetTeamMember(ctx, input.MemberUserID, input.TeamID)

	if err != nil {
		return output, fmt.Errorf("s.repository.GetTeamMember: %w", err)
	}

	if invitedMember.Role != domain.RoleOwner && invitedMember.Role != domain.RoleAdmin {
		return output, domain.ErrNotAllowed
	}

	member, err := domain.NewTeamMember(input.InvitedUserID, input.TeamID, domain.RoleMember)

	if err != nil {
		return output, fmt.Errorf("domain.NewTeamMember: %w", err)
	}

	err = s.repository.CreateTeamMember(ctx, &member)

	if err != nil {
		return output, fmt.Errorf("s.repository.CreateTeamMember: %w", err)
	}

	return dto.InviteOutput{
		TeamMember: member,
	}, nil
}
