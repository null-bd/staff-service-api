package staff

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/null-bd/logger"
)

// region Definition

type (
	IStaffRepository interface {
	}

	staffRepository struct {
		db  *pgxpool.Pool
		log logger.Logger
	}
)

func NewStaffRepository(db *pgxpool.Pool, logger logger.Logger) IStaffRepository {
	return &staffRepository{
		db:  db,
		log: logger,
	}
}

// region SQL Queries
