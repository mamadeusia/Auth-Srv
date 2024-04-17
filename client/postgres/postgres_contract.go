package postgres

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"go-micro.dev/v4/logger"
)

type PostgresRepository struct {
	DB      *pgxpool.Pool
	Queries *Queries
}

func NewPostgres(ctx context.Context, url string) (*PostgresRepository, error) {

	db, err := pgxpool.Connect(ctx, url)
	if err != nil {
		logger.Fatal(err)
		return nil, err
	}

	err = db.Ping(ctx)
	if err != nil {
		if err != nil {
			logger.Fatal(err)
			return nil, err
		}
	}

	pr := &PostgresRepository{
		DB:      db,
		Queries: New(db),
	}

	return pr, nil
}
