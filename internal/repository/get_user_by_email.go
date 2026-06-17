package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/ilydyu/task_manager.git/internal/domain"
	"github.com/ilydyu/task_manager.git/pkg/transaction"
)

func (r *Repository) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	const query = `select id, name, password_hash, created_at from users where email = ?`

	tx := transaction.TryExtractTX(ctx)

	row := tx.QueryRowContext(ctx, query, email)

	user := domain.User{
		Email: email,
	}

	err := row.Scan(&user.ID, &user.Name, &user.PasswordHash, &user.CreatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, domain.ErrNotFound
		}
		return user, fmt.Errorf("row.Scan: %w", err)
	}

	return user, nil
}
