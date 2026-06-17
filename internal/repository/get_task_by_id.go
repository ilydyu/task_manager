package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/ilydyu/task_manager.git/internal/domain"
	"github.com/ilydyu/task_manager.git/pkg/transaction"
)

func (r *Repository) GetTaskByID(ctx context.Context, id int64) (domain.Task, error) {
	const query = `select * from tasks where id = ?`

	tx := transaction.TryExtractTX(ctx)

	row := tx.QueryRowContext(ctx, query, id)

	var task domain.Task

	err := row.Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Status,
		&task.Priority,
		&task.TeamID,
		&task.AssigneeID,
		&task.CreatedBy,
		&task.Deadline,
		&task.CreatedAt,
		&task.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return task, domain.ErrNotFound
		}
		return task, fmt.Errorf("row.Scan: %w", err)
	}

	return task, nil
}
