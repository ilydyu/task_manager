package repository

import (
	"context"
	"fmt"

	"github.com/ilydyu/task_manager.git/pkg/transaction"
)

func (r *Repository) IsTeamMemberExists(ctx context.Context, userID, teamID int64) (bool, error) {
	const query = `select exists (select 1 from team_members where user_id = ? and team_id = ?)`

	tx := transaction.TryExtractTX(ctx)

	var exists bool

	err := tx.QueryRowContext(ctx, query, userID, teamID).Scan(&exists)

	if err != nil {
		return false, fmt.Errorf("tx.QueryRowContext: %w", err)
	}

	return exists, nil
}
