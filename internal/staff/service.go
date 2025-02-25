package staff

import (
	"github.com/null-bd/logger"
)

type IStaffService interface {
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
