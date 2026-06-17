package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/ilydyu/task_manager.git/internal/domain"
	"github.com/ilydyu/task_manager.git/pkg/transaction"
)

func (r *Repository) GetTeamMember(ctx context.Context, userID, teamID int64) (domain.TeamMember, error) {
	const query = `select id, role, created_at, updated_at from team_members where user_id = ? and team_id = ?`

	tx := transaction.TryExtractTX(ctx)

	row := tx.QueryRowContext(ctx, query, userID, teamID)

	member := domain.TeamMember{
		UserID: userID,
		TeamID: teamID,
	}

	err := row.Scan(&member.ID, &member.Role, &member.CreatedAt, &member.UpdatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return member, domain.ErrNotFound
		}
		return member, fmt.Errorf("row.Scan: %w", err)
	}

	return member, nil
}
