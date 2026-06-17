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

func (r *Repository) CreateTeam(ctx context.Context, team *domain.Team) error {
	const query = `insert into teams (name, created_by) values (?, ?)`

	tx := transaction.TryExtractTX(ctx)

	res, err := tx.ExecContext(ctx, query, team.Name, team.CreatedBy)

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

	team.ID = id

	return nil
}
