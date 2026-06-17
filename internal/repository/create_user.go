package repository

import (
	"context"
	"fmt"

	"github.com/ilydyu/task_manager.git/internal/domain"
	"github.com/ilydyu/task_manager.git/pkg/transaction"
)

func (r *Repository) CreateUser(ctx context.Context, user *domain.User) error {
	const query = `insert into users (name, email, password_hash) values (?, ?, ?)`

	tx := transaction.TryExtractTX(ctx)

	res, err := tx.ExecContext(ctx, query, user.Name, user.Email, user.PasswordHash)

	if err != nil {
		return fmt.Errorf("tx.ExecContext: %w", err)
	}

	id, err := res.LastInsertId()

	if err != nil {
		return fmt.Errorf("res.LastInsertId: %w", err)
	}

	user.ID = id

	return nil
}
