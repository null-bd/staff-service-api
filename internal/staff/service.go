package staff

import (
	"context"
	stderr "errors"

	"github.com/google/uuid"
	"github.com/null-bd/logger"
	"github.com/null-bd/staff-service-api/internal/errors"
)

type IStaffService interface {
	CreateStaff(ctx context.Context, staff *Staff) (*Staff, error)
}

type staffService struct {
	repo IStaffRepository
	log  logger.Logger
}

func NewStaffService(repo IStaffRepository, logger logger.Logger) IStaffService {
	return &staffService{
		repo: repo,
		log:  logger,
	}
}

func (s *staffService) CreateStaff(ctx context.Context, staff *Staff) (*Staff, error) {
	s.log.Info("service : CreateStaff : begin", nil)

	existingStaff, err := s.repo.GetbyID(ctx, staff.ID)
	if err != nil {
		return nil, err
	}
	if existingStaff != nil {
		return nil, &errors.AppError{
			Code:    errors.ErrStaffExists,
			Message: "staff with this code already exists",
			Err:     stderr.New("organization with this code already exists"),
		}
	}

	staff.ID = uuid.New().String()
	staff.Status = "inactive"

	createdStaff, err := s.repo.Create(ctx, staff)
	if err != nil {
		return nil, err
	}

	s.log.Info("service: CreatedStaff: exit", nil)
	return createdStaff, nil
}
