package query

import (
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	"github.com/georgysavva/scany/sqlscan"
	"github.com/jackc/pgx/v4"
	"github.com/rs/zerolog/log"
)

func Builder() goqu.DialectWrapper {
	goqu.SetDefaultPrepared(true)
	return goqu.Dialect("postgres")
}

func Scan[T any](ctx context.Context, conn *sql.DB, b *goqu.SelectDataset) ([]*T, error) {
	q, args, err := b.ToSQL()
	if err != nil {
		return nil, err
	}

	rows, err := conn.QueryContext(ctx, q, args...)
	if err != nil && err != pgx.ErrNoRows {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		closeErr := rows.Close()
		if closeErr != nil {
			log.Err(err).Send()
		}
	}(rows)

	var res []*T
	err = sqlscan.ScanAll(&res, rows)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func ScanOne[T any](ctx context.Context, conn *sql.DB, b *goqu.SelectDataset) (*T, error) {
	all, err := Scan[T](ctx, conn, b)
	if err != nil {
		return nil, err
	}

	if all == nil || len(all) == 0 {
		return nil, nil
	}

	return all[0], nil
}

type ExecAble interface {
	ToSQL() (string, []interface{}, error)
}

func Exec(ctx context.Context, conn *sql.DB, b ExecAble) error {
	q, args, err := b.ToSQL()
	if err != nil {
		return err
	}

	_, err = conn.ExecContext(ctx, q, args...)
	return err
}
