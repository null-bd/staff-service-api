package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/null-bd/logger"
	"github.com/null-bd/microservice-name/internal/health"
)

// region Definition
type (
	IHealthHandler interface {
		HealthCheck(c *gin.Context)
	}
	healthHandler struct {
		healthSvc health.IHealthService
		log       logger.Logger
	}
)

func NewHealthHandler(healthSvc health.IHealthService, logger logger.Logger) IHealthHandler {
	return &healthHandler{
		healthSvc: healthSvc,
		log:       logger,
	}
}

// region Implementation

func (h *healthHandler) HealthCheck(c *gin.Context) {
	h.log.Info("handler : HealthCheck : begin", nil)
	status, err := h.healthSvc.CheckHealth()
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"details": status,
	})
	h.log.Info("handler : HealthCheck : exit", nil)
}
