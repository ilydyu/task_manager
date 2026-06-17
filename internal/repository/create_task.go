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

func (r *Repository) CreateTask(ctx context.Context, task *domain.Task) error {
	const query = `insert into tasks
	(title, description, status, priority, team_id, assignee_id, created_by, deadline) 
	values 
	(?, ?, ?, ?, ?, ?, ?, ?)`

	tx := transaction.TryExtractTX(ctx)

	res, err := tx.ExecContext(
		ctx,
		query,
		task.Title,
		task.Description,
		task.Status,
		task.Priority,
		task.TeamID,
		task.AssigneeID,
		task.CreatedBy,
		task.Deadline,
	)

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

	task.ID = id

	return nil
}
