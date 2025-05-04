package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/null-bd/staff-service-api/internal/staff"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockStaffSvc struct {
	mock.Mock
}

func (m *mockStaffSvc) CreateStaff(ctx context.Context, s *staff.Staff) (*staff.Staff, error) {
	args := m.Called(ctx, s)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*staff.Staff), args.Error(1)
}

func setupTest(t *testing.T) (*gin.Engine, *mockStaffSvc, *mockLogger) {
	t.Log("Setting up test")
	gin.SetMode(gin.TestMode)

	mockStaffSvc := new(mockStaffSvc)
	mockLog := new(mockLogger)

	handler := NewStaffHandler(mockStaffSvc, mockLog)

	router := gin.New()
	router.POST("/staffs", handler.CreateStaff)

	return router, mockStaffSvc, mockLog
}

func TestCreateStaff(t *testing.T) {
	router, mockStaffSvc, mockLog := setupTest(t)

	mockLog.On("Info", "handler : CreateStaff : begin", mock.Anything).Return()
	mockLog.On("Info", "handler : CreateStaff : exit", mock.Anything).Return()

	inputDTO := CreateStaffRequest{
		FirstName: "Test Staff",
		Code:      "TEST001",
		StaffType: "doctor",
	}

	expectedStaff := &staff.Staff{
		ID:        "test-id-1",
		FirstName: "Test Staff",
		Code:      "TEST001",
		StaffType: "doctor",
		Status:    "inactive",
	}

	mockStaffSvc.On("CreateStaff", mock.Anything, mock.MatchedBy(func(staff *staff.Staff) bool {
		return staff.FirstName == inputDTO.FirstName && staff.Code == inputDTO.Code
	})).Return(expectedStaff, nil)

	body, _ := json.Marshal(inputDTO)
	req, _ := http.NewRequest("POST", "/staffs", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response StaffResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, expectedStaff.ID, response.ID)
	assert.Equal(t, inputDTO.FirstName, response.FirstName)

	mockStaffSvc.AssertExpectations(t)
	mockLog.AssertExpectations(t)
}
