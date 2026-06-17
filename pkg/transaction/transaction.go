package transaction

import (
	"context"
	"database/sql"

	"github.com/ilydyu/task_manager.git/pkg/mysql"
)

var (
	pool *mysql.Pool
)

type ctxKey struct{}

func Init(p *mysql.Pool) {
	pool = p
}

type Executor interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
}

func TryExtractTX(ctx context.Context) Executor { //nolint:ireturn
	tx, ok := ctx.Value(ctxKey{}).(*sql.Tx)
	if !ok {
		return pool
	}

	return tx
}
