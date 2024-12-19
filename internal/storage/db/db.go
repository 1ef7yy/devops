package db

import (
	"context"
	"sync"

	"github.com/1ef7yy/devops-course/pkg/logger"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	Logger   logger.Logger
	Database *pgxpool.Pool
}

func NewPostgres(ctx context.Context, dsn string, log logger.Logger) (*Postgres, error) {
	var (
		pgInstance *Postgres
		pgOnce     sync.Once
		PgErr      error
	)

	pgOnce.Do(func() {
		db, err := pgxpool.New(ctx, dsn)
		if err != nil {
			log.Error("Unable to connect to database: " + err.Error())
			PgErr = err
		}

		pgInstance = &Postgres{
			Logger:   log,
			Database: db,
		}
	})

	if PgErr != nil {
		return nil, PgErr
	}
	return pgInstance, nil
}

func (pg *Postgres) Close() {
	pg.Database.Close()
}
