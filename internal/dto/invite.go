package dto

import "github.com/ilydyu/task_manager.git/internal/domain"

type InviteInput struct {
	MemberUserID  int64 `json:"member_user_id"`
	InvitedUserID int64 `json:"invited_user_id"`
	TeamID        int64 `json:"team_id"`
}

type InviteOutput struct {
	domain.TeamMember
}
