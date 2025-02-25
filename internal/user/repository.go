package user

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/null-bd/logger"
)

// region Definition

type (
	IUserRepository interface {
	}

	userRepository struct {
		db  *pgxpool.Pool
		log logger.Logger
	}
)

func NewUserRepository(db *pgxpool.Pool, logger logger.Logger) IUserRepository {
	return &userRepository{
		db:  db,
		log: logger,
	}
}

// region SQL Queries
