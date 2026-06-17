package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/ilydyu/task_manager.git/internal/domain"
	"github.com/ilydyu/task_manager.git/pkg/transaction"
)

func (r *Repository) UpdateTask(ctx context.Context, task *domain.Task) error {
	const query = `update tasks set title=?, description=?, status=?, priority=?, assignee_id=?, deadline=? where id = ?`

	tx := transaction.TryExtractTX(ctx)

	res, err := tx.ExecContext(
		ctx,
		query,
		task.Title,
		task.Description,
		task.Status,
		task.Priority,
		task.AssigneeID,
		task.Deadline,
		task.ID,
	)

	if err != nil {
		if strings.Contains(err.Error(), "foreign key constraint") ||
			strings.Contains(err.Error(), "violates foreign key constraint") {
			return domain.ErrNotFound
		}
		return fmt.Errorf("tx.ExecContext: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("res.RowsAffected: %w", err)
	}

	if rowsAffected == 0 {
		return domain.ErrNotFound
	}

	return nil
}
