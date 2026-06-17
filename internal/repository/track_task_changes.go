package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/ilydyu/task_manager.git/internal/domain"
	"github.com/ilydyu/task_manager.git/pkg/transaction"
)

func (r *Repository) TrackTaskChanges(ctx context.Context, task domain.Task, action string, oldValue []byte) error {
	if action != "create" && action != "update" {
		return domain.ErrInvalidAction
	}

	query := `
		insert into task_history 
		(task_id, changed_by, field_name, old_value, new_value, action, changed_at)
		values (?, ?, ?, ?, ?, ?, ?)
	`

	tx := transaction.TryExtractTX(ctx)

	data, err := json.Marshal(task)
	if err != nil {
		return fmt.Errorf("json.Marshal: %w", err)
	}

	_, err = tx.ExecContext(
		ctx,
		query,
		task.ID,
		task.CreatedBy,
		"all_fields",
		oldValue,
		string(data),
		action,
		time.Now(),
	)

	if err != nil {
		return fmt.Errorf("tx.ExecContext: %w", err)
	}

	return nil
}
