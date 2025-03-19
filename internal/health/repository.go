package health

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/null-bd/logger"
	"github.com/null-bd/staff-service-api/internal/errors"
)

type (
	iHealthRepository interface {
		CheckDatabase() error
	}

	healthRepository struct {
		db  *pgxpool.Pool
		log logger.Logger
	}
)

func NewHealthRepository(db *pgxpool.Pool, logger logger.Logger) iHealthRepository {
	return &healthRepository{
		db:  db,
		log: logger,
	}
}

func (r *healthRepository) CheckDatabase() error {
	r.log.Debug("repository : CheckDatabase : begin", nil)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := r.db.Ping(ctx); err != nil {
		return errors.NewDatabaseConnectionError(err)
	}
	r.log.Debug("repository : CheckDatabase : exit", nil)
	return nil
}
