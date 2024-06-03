package container

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	. "github.com/mariemalysheva/tokenized-reinvestment"
	"github.com/mariemalysheva/tokenized-reinvestment/config"
	"github.com/pressly/goose/v3"
)

func (c *Container) GetDB(ctx context.Context) (*pgxpool.Pool, error) {
	if c.db != nil {
		return c.db, nil
	}

	dbConn := config.DBConn

	pool, err := pgxpool.New(ctx, dbConn)
	if err != nil {
		return nil, err
	}
	c.db = pool

	connConfig, err := pgx.ParseConfig(dbConn)
	if err != nil {
		return nil, err
	}
	stdlib.RegisterConnConfig(connConfig)

	goose.SetBaseFS(EmbedMigrations)

	gooseDB, err := goose.OpenDBWithDriver("pgx", dbConn)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := gooseDB.Close(); err != nil {
			fmt.Println("failed to close goose db", "err", err.Error())
		}
	}()

	// Apply migrations on app startup
	if err := goose.Up(gooseDB, "migrations", goose.WithAllowMissing()); err != nil {
		return nil, err
	}

	return c.db, nil
}
