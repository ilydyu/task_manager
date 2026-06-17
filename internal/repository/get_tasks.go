package repository

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ilydyu/task_manager.git/internal/domain"
	"github.com/ilydyu/task_manager.git/pkg/transaction"
)

func (r *Repository) GetTasks(ctx context.Context, teamID, assigneeID, status, cursor string) ([]domain.Task, error) {
	query := `SELECT * FROM tasks WHERE 1=1`
	args := []any{}
	tasks := []domain.Task{}

	if teamID != "" {
		query += " AND team_id = ?"
		args = append(args, teamID)
	}

	if assigneeID != "" {
		query += " AND assignee_id = ?"
		args = append(args, assigneeID)
	}

	if status != "" {
		query += " AND status = ?"
		args = append(args, status)
	}

	if cursor != "" {
		_, err := strconv.Atoi(cursor)

		if err != nil {
			return tasks, fmt.Errorf("repository: GetTasks: strconv.Atoi %w", err)
		}

		query += " AND id > ?"
		args = append(args, cursor)
	}

	query += " ORDER BY id ASC LIMIT 21"

	tx := transaction.TryExtractTX(ctx)

	rows, err := tx.QueryContext(ctx, query, args...)

	if err != nil {
		return tasks, fmt.Errorf("tx.QueryContext: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var t domain.Task

		err := rows.Scan(
			&t.ID,
			&t.Title,
			&t.Description,
			&t.Status,
			&t.Priority,
			&t.TeamID,
			&t.AssigneeID,
			&t.CreatedBy,
			&t.Deadline,
			&t.CreatedAt,
			&t.UpdatedAt,
		)

		if err != nil {
			return tasks, fmt.Errorf("row.Scan: %w", err)
		}

		tasks = append(tasks, t)
	}

	if rows.Err() != nil {
		return tasks, fmt.Errorf("rows.Err: %w", err)
	}

	return tasks, nil
}
