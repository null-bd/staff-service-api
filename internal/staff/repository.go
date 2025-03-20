package staff

import (
	"context"
	stderr "errors"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/null-bd/logger"
	"github.com/null-bd/staff-service-api/internal/errors"
)

// region Definition

type (
	IStaffRepository interface {
		Create(ctx context.Context, staff *Staff) (*Staff, error)
		GetbyID(ctx context.Context, id string) (*Staff, error)
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
			id, branch_id, organization_id, first_name, last_name, code, status, type, specialties, 
			departments_Id, departments_role, departments_isprimary, schedule_type, schedule_shifts, 
			email, phone, date_of_birth, gender, address,
			metadata, created_at, updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, 
			$7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21
		) RETURNING id, created_at`

	getDeptByCodeQuery = `
		SELECT 
		id, branch_id, organization_id, first_name, last_name, code, status, type, specialties, 
			departments_Id, departments_role, departments_isprimary, schedule_type, schedule_shifts, 
			email, phone, date_of_birth, gender, address,
			metadata, created_at, updated_at
		FROM staffs
		WHERE code = $1 AND deleted_at IS NULL`
)

func (r *staffRepository) Create(ctx context.Context, staff *Staff) (*Staff, error) {
	r.log.Debug("repository : Create : begin", nil)
	now := time.Now().UTC()

	// Execute query
	_, err := r.db.Exec(ctx, createStaffQuery,
		&staff.ID,
		&staff.BranchID,
		&staff.OrganizationID,
		&staff.FirstName,
		&staff.LastName,
		&staff.Code,
		&staff.Status,
		&staff.StaffType,
		&staff.Specialities,
		&staff.Departments.DepartmentID,
		&staff.Departments.Role,
		&staff.Departments.IsPrimary,
		&staff.Schedule.Type,
		&staff.Schedule.Shifts,
		&staff.Email,
		&staff.Phone,
		&staff.DateOfBirth,
		&staff.Gender,
		&staff.Address,
		&staff.Metadata,
		now.Format(time.RFC3339),
	)
	if err != nil {
		return nil, errors.New(errors.ErrDatabaseOperation, "database error", err)
	}

	createdStaff, err := r.GetbyID(ctx, staff.ID)
	if err != nil {
		return nil, err
	}

	r.log.Debug("repository : Create : exit", nil)
	return createdStaff, nil
}

func (r *staffRepository) GetbyID(ctx context.Context, id string) (*Staff, error) {
	r.log.Debug("repository : GetbyID : begin", nil)

	staff := &Staff{
		Departments: Departments{},
		Schedule:    Schedule{},
		Address:     Address{},
		Metadata:    make(map[string]interface{}),
	}

	var createdAt, updatedAt time.Time

	err := r.db.QueryRow(ctx, getDeptByCodeQuery, id).Scan(
		&staff.ID,
		&staff.BranchID,
		&staff.OrganizationID,
		&staff.FirstName,
		&staff.LastName,
		&staff.Code,
		&staff.Status,
		&staff.StaffType,
		&staff.Specialities,
		&staff.Departments.DepartmentID,
		&staff.Departments.Role,
		&staff.Departments.IsPrimary,
		&staff.Schedule.Type,
		&staff.Schedule.Shifts,
		&staff.Email,
		&staff.Phone,
		&staff.DateOfBirth,
		&staff.Gender,
		&staff.Address,
		&staff.Metadata,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		if stderr.Is(err, pgx.ErrNoRows) {
			return nil, errors.New(errors.ErrStaffNotFound, "Staff not found", err)
		}
		return nil, errors.New(errors.ErrDatabaseOperation, "database error", err)
	}

	staff.CreatedAt = createdAt.Format(time.RFC3339)
	staff.UpdatedAt = updatedAt.Format(time.RFC3339)

	r.log.Debug("repository : GetByID : exit", logger.Fields{"department": staff})
	return staff, nil
}
