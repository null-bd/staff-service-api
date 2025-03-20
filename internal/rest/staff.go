package rest

import (
	"github.com/null-bd/logger"
	"github.com/null-bd/staff-service-api/internal/staff"
)

type IStaffHandler interface {
}

type staffHandler struct {
	staffSvc staff.IStaffService
	log      logger.Logger
}

func NewStaffHandler(staffSvc staff.IStaffService, logger logger.Logger) IStaffHandler {
	return &staffHandler{
		staffSvc: staffSvc,
		log:      logger,
	}
}
