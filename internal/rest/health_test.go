package rest

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/null-bd/microservice-name/internal/errors"
	"github.com/null-bd/microservice-name/internal/health"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockHealthService struct {
	mock.Mock
}

func (m *mockHealthService) CheckHealth() (*health.HealthStatus, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*health.HealthStatus), args.Error(1)
}

// Test setup helper
func setupHealthTest(t *testing.T) (*gin.Engine, *mockHealthService, *mockLogger) {
	t.Log("Setting up test...")
	gin.SetMode(gin.TestMode)

	mockHealthSvc := new(mockHealthService)
	mockLog := new(mockLogger)

	// Setup logging expectations
	mockLog.On("Info", "handler : HealthCheck : begin", mock.Anything).Return()

	handler := &healthHandler{
		healthSvc: mockHealthSvc,
		log:       mockLog,
	}

	router := gin.New()
	router.Use(gin.Recovery())
	router.GET("/health", handler.HealthCheck)

	return router, mockHealthSvc, mockLog
}

func TestHealthCheck_Success(t *testing.T) {
	// Arrange
	router, mockHealthSvc, mockLog := setupHealthTest(t)

	expectedStatus := &health.HealthStatus{
		Database: health.HealthComponent{
			Status:  "healthy",
			Message: "Connected",
		},
	}

	mockHealthSvc.On("CheckHealth").Return(expectedStatus, nil)
	mockLog.On("Info", "handler : HealthCheck : exit", mock.Anything).Return()

	// Act
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "healthy", response["status"])

	mockHealthSvc.AssertExpectations(t)
	mockLog.AssertExpectations(t)
}

func TestHealthCheck_ServiceError(t *testing.T) {
	// Arrange
	router, mockHealthSvc, mockLog := setupHealthTest(t)

	expectedError := errors.New(errors.ErrDatabaseConnection, "Database connection failed", nil)
	mockHealthSvc.On("CheckHealth").Return(nil, expectedError)

	// Act
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusServiceUnavailable, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, string(errors.ErrDatabaseConnection), response["code"])

	mockHealthSvc.AssertExpectations(t)
	mockLog.AssertExpectations(t)
}

func TestHealthCheck_PanicRecovery(t *testing.T) {
	// Arrange
	router, mockHealthSvc, mockLog := setupHealthTest(t)

	mockHealthSvc.On("CheckHealth").Run(func(args mock.Arguments) {
		panic("test panic")
	})

	// Act
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockLog.AssertExpectations(t)
}
