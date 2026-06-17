package transaction

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/rs/zerolog/log"
)

func Wrap(ctx context.Context, fn func(context.Context) error) error {
	tx, err := pool.Begin()

	if err != nil {
		return fmt.Errorf("pool.Begin: %w", err)
	}

	defer func() {
		err = tx.Rollback()
		if err != nil && !errors.Is(err, sql.ErrTxDone) {
			log.Error().Err(err).Msg("transaction: Rollback")
		}
	}()

	ctx = context.WithValue(ctx, ctxKey{}, tx)

	err = fn(ctx)
	if err != nil {
		return fmt.Errorf("fn: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("tx.Commit: %w", err)
	}

	return nil
}
