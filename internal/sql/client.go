package sql

import (
	"context"
	"database/sql"
	"fmt"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/kostaspt/go-tiny/config"
	"github.com/kostaspt/go-tiny/internal/sql/ent"
	"github.com/kostaspt/go-tiny/internal/sql/ent/migrate"
)

func NewClient(ctx context.Context, c *config.Config) (client *ent.Client, cleanup func(), err error) {
	db, err := sql.Open("pgx", fmt.Sprintf("postgresql://%s:%s@%s/%s", c.SQL.Username, c.SQL.Password, c.SQL.Addr, c.SQL.Database))
	if err != nil {
		return
	}

	cleanup = func() {
		db.Close()
	}

	drv := entsql.OpenDB(dialect.Postgres, db)
	client = ent.NewClient(ent.Driver(drv))

	err = client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	return
}
