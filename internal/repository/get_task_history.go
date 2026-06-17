package repository

import (
	"context"
	"fmt"

	"github.com/ilydyu/task_manager.git/internal/domain"
	"github.com/ilydyu/task_manager.git/pkg/transaction"
)

func (r *Repository) GetTaskHistory(ctx context.Context, taskID int64) ([]domain.TaskHistory, error) {
	const query = `select * from task_history where task_id = ?`
	history := []domain.TaskHistory{}

	tx := transaction.TryExtractTX(ctx)

	rows, err := tx.QueryContext(ctx, query, taskID)

	if err != nil {
		return history, fmt.Errorf("tx.QueryContext: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var th domain.TaskHistory

		err := rows.Scan(&th.ID, &th.TaskID, &th.ChangedBy, &th.FieldName, &th.Action, &th.OldValue, &th.NewValue, &th.ChangedAt)

		if err != nil {
			return history, fmt.Errorf("row.Scan: %w", err)
		}

		history = append(history, th)
	}

	if rows.Err() != nil {
		return history, fmt.Errorf("rows.Err: %w", err)
	}

	return history, nil
}
