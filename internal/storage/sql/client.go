package sql

import (
	"context"
	"database/sql"
	"embed"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/pgx"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/rs/zerolog/log"

	"github.com/kostaspt/go-tiny/config"
)

//go:embed migrations/*.sql
var fs embed.FS

type DB struct {
	Conn   *sql.DB
	config *config.Config
}

func NewConnection(ctx context.Context, c *config.Config) (*DB, func(), error) {
	var err error

	s := &DB{config: c}

	s.Conn, err = sql.Open("pgx", s.dsn(c.SQL.Database))
	if err != nil {
		return nil, nil, err
	}

	if err = s.VerifyDatabase(ctx); err != nil {
		return nil, nil, err
	}

	if err = s.Migrate(); err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		if err = s.Conn.Close(); err != nil {
			log.Err(err).Send()
			return
		}
	}

	return s, cleanup, nil
}

func (db *DB) VerifyDatabase(ctx context.Context) error {
	conn, err := sql.Open("pgx", db.dsn("postgres"))
	if err != nil {
		return err
	}
	defer conn.Close()

	row := conn.QueryRowContext(ctx, `SELECT COUNT(*) FROM pg_database WHERE datname = '`+db.config.SQL.Database+`'`)
	if row.Err() != nil {
		return row.Err()
	}

	var cnt uint16
	if err = row.Scan(&cnt); err != nil {
		return err
	}

	// Database exist, all good
	if cnt == 1 {
		return nil
	}

	_, err = conn.Exec(`CREATE DATABASE ` + db.config.SQL.Database)
	return err
}

func (db *DB) Migrate() error {
	f, err := iofs.New(fs, "migrations")
	if err != nil {
		return err
	}

	c, err := pgx.WithInstance(db.Conn, &pgx.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithInstance("iofs", f, db.config.SQL.Database, c)
	if err != nil {
		return err
	}

	if err = m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			log.Info().Msg("No migrations to run.")
		} else {
			return err
		}
	}

	return nil
}

func (db *DB) dsn(database string) string {
	return fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", db.config.SQL.Username, db.config.SQL.Password, db.config.SQL.Addr, database)
}
