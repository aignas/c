package dbinsert

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
)

const (
	_table = "mytable"
)

type Row struct {
	ID   uuid.UUID `db:"id"`
	Name string    `db:"name"`
}

func InsertCached(ctx context.Context, db *sqlx.DB, data []Row) error {
	return WithTransaction(ctx, db, func(ctx context.Context, tx sqlx.PreparerContext) error {
		for _, d := range data {
			q := squirrel.Insert(_table).
				Columns("id", "name").
				Values(d.ID, d.Name)

			query, args, err := q.ToSql()
			if err != nil {
				return fmt.Errorf("to SQL: %w", err)
			}

			stmt, err := tx.PrepareContext(ctx, query)
			if err != nil {
				return fmt.Errorf("prep: %w", err)
			}

			_, err = stmt.ExecContext(ctx, args...)
			if err != nil {
				return fmt.Errorf("exec: %w", err)
			}
		}
		return nil
	})
}

func InsertBulk(ctx context.Context, db *sqlx.DB, data []Row) error {
	return WithTransaction(ctx, db, func(ctx context.Context, tx sqlx.PreparerContext) error {
		q := squirrel.Insert(_table).
			Columns("id", "name")

		for i := range data {
			q = q.Values(data[i].ID, data[i].Name)
		}

		query, args, err := q.ToSql()
		if err != nil {
			return fmt.Errorf("to SQL: %w", err)
		}

		stmt, err := tx.PrepareContext(ctx, query)
		if err != nil {
			return fmt.Errorf("prep: %w", err)
		}

		_, err = stmt.ExecContext(ctx, args...)
		if err != nil {
			return fmt.Errorf("exec: %w", err)
		}
		return nil
	})
}

func WithTransaction(ctx context.Context, db *sqlx.DB, f func(ctx context.Context, tx sqlx.PreparerContext) error) (retErr error) {
	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if retErr == nil {
			if err := tx.Commit(); err != nil {
				retErr = fmt.Errorf("commit: %w", err)
			}
			return
		}

		if err := tx.Rollback(); err != nil {
			retErr = fmt.Errorf("%w; rollback: %w", retErr, err)
		}
	}()

	err = f(ctx, squirrel.NewStmtCache(tx))
	if err != nil {
		return fmt.Errorf("func: %w", err)
	}

	_, err = tx.ExecContext(ctx, "DELETE FROM mytable")
	return err
}
