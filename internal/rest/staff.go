package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/null-bd/logger"
	"github.com/null-bd/staff-service-api/internal/errors"
	"github.com/null-bd/staff-service-api/internal/staff"
)

type IStaffHandler interface {
	CreateStaff(c *gin.Context)
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

func (h *staffHandler) CreateStaff(c *gin.Context) {
	h.log.Info("handler : CreateStaff : begin", nil)

	var req CreateStaffRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		HandleError(c, errors.New(errors.ErrBadRequest, "invalid request body", err))
		return
	}

	staff := ToStaff(&req)
	result, err := h.staffSvc.CreateStaff(c.Request.Context(), staff)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, ToStaffResponse(result))
	h.log.Info("handler : CreateStaff : exit", nil)
}
