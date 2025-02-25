package rest

import (
	"net/http"

	"github.com/null-bd/microservice-name/internal/errors"

	"github.com/gin-gonic/gin"
)

// ErrorResponse represents the API error response structure
type ErrorResponse struct {
	Code    string               `json:"code"`
	Message string               `json:"message"`
	Details []errors.ErrorDetail `json:"details,omitempty"`
}

// HandleError converts application errors to HTTP responses
func HandleError(c *gin.Context, err error) {
	if err == nil {
		return
	}

	// Check if it's our custom error type
	if appErr, ok := err.(*errors.AppError); ok {
		status := getHTTPStatusFromErrorCode(appErr.Code)
		c.JSON(status, ErrorResponse{
			Code:    string(appErr.Code),
			Message: appErr.Message,
			Details: appErr.Details,
		})
		return
	}

	// Handle unexpected errors
	c.JSON(http.StatusInternalServerError, ErrorResponse{
		Code:    "INTERNAL_ERROR",
		Message: "An unexpected error occurred",
	})
}

func getHTTPStatusFromErrorCode(code errors.ErrorCode) int {
	switch code {
	case errors.ErrDatabaseConnection, errors.ErrDatabaseQuery:
		return http.StatusServiceUnavailable
	case errors.ErrBadRequest:
		return http.StatusBadRequest
	case errors.ErrStaffNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
