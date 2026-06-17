package domain

import (
	"fmt"
	"time"
)

type MemberRole string

const (
	RoleOwner  MemberRole = "owner"
	RoleAdmin  MemberRole = "admin"
	RoleMember MemberRole = "member"
)

type TeamMember struct {
	ID        int64      `json:"id"`
	TeamID    int64      `json:"team_id" validate:"required"`
	UserID    int64      `json:"user_id" validate:"required"`
	Role      MemberRole `json:"role" validate:"required,oneof=owner admin member"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

func NewTeamMember(userID, teamID int64, role MemberRole) (TeamMember, error) {
	switch role {
	case RoleOwner, RoleAdmin, RoleMember:
	default:
		return TeamMember{}, fmt.Errorf("invalid role: %s", role)
	}

	m := TeamMember{
		UserID: userID,
		TeamID: teamID,
		Role:   role,
	}

	if err := m.Validate(); err != nil {
		return TeamMember{}, fmt.Errorf("p.Validate: %w", err)
	}

	return m, nil
}

func (m TeamMember) Validate() error {
	err := validate.Struct(m)
	if err != nil {
		return fmt.Errorf("validate.Struct: %w", err)
	}

	return nil
}
