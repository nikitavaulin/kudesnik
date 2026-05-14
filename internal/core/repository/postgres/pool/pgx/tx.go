package core_pgx_pool

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	core_postgres_pool "github.com/nikitavaulin/kudesnik/internal/core/repository/postgres/pool"
)

type Tx struct {
	tx pgx.Tx
}

func (t *Tx) Exec(ctx context.Context, sql string, arguments ...any) (core_postgres_pool.CommandTag, error) {
	cmdTag, err := t.tx.Exec(ctx, sql, arguments...)
	if err != nil {
		return nil, err
	}
	return pgxCommandTag{cmdTag}, nil
}

func (t *Tx) Query(ctx context.Context, sql string, args ...any) (core_postgres_pool.Rows, error) {
	rows, err := t.tx.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	return pgxRows{rows}, nil
}

func (t *Tx) QueryRow(ctx context.Context, sql string, args ...any) core_postgres_pool.Row {
	row := t.tx.QueryRow(ctx, sql, args...)
	return pgxRow{row}
}

func (t *Tx) Commit(ctx context.Context) error {
	if err := t.tx.Commit(ctx); err != nil {
		return fmt.Errorf("commit transaction: %w", err)
	}
	return nil
}

func (t *Tx) Rollback(ctx context.Context) error {
	if err := t.tx.Rollback(ctx); err != nil {
		return fmt.Errorf("rollback transaction: %w", err)
	}
	return nil
}
