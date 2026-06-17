package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/ilydyu/task_manager.git/internal/domain"
	"github.com/ilydyu/task_manager.git/pkg/transaction"
)

func (r *Repository) CreateTeamMember(ctx context.Context, member *domain.TeamMember) error {
	const query = `insert into team_members (user_id, team_id, role) values (?, ?, ?)`

	tx := transaction.TryExtractTX(ctx)

	res, err := tx.ExecContext(ctx, query, member.UserID, member.TeamID, member.Role)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) || strings.Contains(err.Error(), "foreign key constraint") ||
			strings.Contains(err.Error(), "violates foreign key constraint") {
			return domain.ErrNotFound
		}
		return fmt.Errorf("tx.ExecContext: %w", err)
	}

	id, err := res.LastInsertId()

	if err != nil {
		return fmt.Errorf("res.LastInsertId: %w", err)
	}

	member.ID = id

	return nil
}
