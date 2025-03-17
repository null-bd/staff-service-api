package staff

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/null-bd/logger"
)

// region Definition

type (
	IStaffRepository interface {
		Create(ctx context.Context, dept *Staff) (*Staff, error)
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
const (
	createStaffQuery = `
	    INSERT INTO staffs (
			id, branch_id, organization_id, name, code, type, specialty, 
			parent_department_id, status, capacity_total_beds, capacity_available_beds, 
			capacity_operating_rooms, operating_hours_weekday, operating_hours_weekend, 
			operating_hours_timezone, operating_hours_holidays, department_head_id,
			metadata, created_at, updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, 
			$7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $19
		) RETURNING id, updated_at`
)

func (r *staffRepository) Create(ctx context.Context, staff *Staff) (*Staff, error) {
	r.log.Debug("repository : Create : begin", nil)

	r.log.Debug("repository : Create : exit", nil)
	return nil, nil
}
