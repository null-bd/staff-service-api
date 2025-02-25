package domain

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/null-bd/logger"
)

// region Definition

type (
	IDomainRepository interface {
	}

	domainRepository struct {
		db  *pgxpool.Pool
		log logger.Logger
	}
)

func NewDomainRepository(db *pgxpool.Pool, logger logger.Logger) IDomainRepository {
	return &domainRepository{
		db:  db,
		log: logger,
	}
}

// region SQL Queries
