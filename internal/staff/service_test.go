package staff

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/null-bd/staff-service-api/internal/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestStaffService_CreateStaff(t *testing.T) {
	tests := []struct {
		name        string
		input       *Staff
		setupMocks  func(*mockRepository, *mockLogger)
		checkResult func(*testing.T, *Staff, error)
	}{
		{
			name: "Success - Create New Staff",
			input: &Staff{
				FirstName: "Test Staff",
				Code:      "TEST001",
				StaffType: "doctor",
			},
			setupMocks: func(repo *mockRepository, logger *mockLogger) {
				// Expect both begin and exit logs for successful case
				logger.On("Info", "service : CreateStaff : begin", mock.Anything).Return()
				logger.On("Info", "service : CreateStaff : exit", mock.Anything).Return()

				repo.On("GetByCode", mock.Anything, "TEST001").Return(nil, nil)
				repo.On("Create", mock.Anything, mock.MatchedBy(func(staff *Staff) bool {
					return staff.Code == "TEST001" && staff.Status == "inactive"
				})).Return(&Staff{
					ID:        uuid.New().String(),
					FirstName: "Test Staff",
					Code:      "TEST001",
					StaffType: "doctor",
					Status:    "inactive",
					CreatedAt: "2024-01-01T00:00:00Z",
					UpdatedAt: "2024-01-01T00:00:00Z",
				}, nil)
			},
			checkResult: func(t *testing.T, result *Staff, err error) {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, "TEST001", result.Code)
				assert.Equal(t, "inactive", result.Status)
				assert.NotEmpty(t, result.ID)
			},
		},
		{
			name: "Error - Staff Already Exists",
			input: &Staff{
				FirstName: "Test Staff",
				Code:      "TEST001",
				StaffType: "doctor",
			},
			setupMocks: func(repo *mockRepository, logger *mockLogger) {
				// Only expect begin log for error case
				logger.On("Info", "service : CreateStaff : begin", mock.Anything).Return()

				existingStaff := &Staff{
					ID:   uuid.New().String(),
					Code: "TEST001",
				}
				repo.On("GetByCode", mock.Anything, "TEST001").Return(existingStaff, nil)
			},
			checkResult: func(t *testing.T, result *Staff, err error) {
				assert.Nil(t, result)
				assert.Error(t, err)
				appErr, ok := err.(*errors.AppError)
				assert.True(t, ok)
				assert.Equal(t, errors.ErrStaffExists, appErr.Code)
			},
		},
		{
			name: "Error - Repository Error on GetByCode",
			input: &Staff{
				FirstName: "Test Staff",
				Code:      "TEST001",
				StaffType: "doctor",
			},
			setupMocks: func(repo *mockRepository, logger *mockLogger) {
				// Only expect begin log for error case
				logger.On("Info", "service : CreateStaff : begin", mock.Anything).Return()

				repo.On("GetByCode", mock.Anything, "TEST001").
					Return(nil, errors.New(errors.ErrDatabaseOperation, "database error", nil))
			},
			checkResult: func(t *testing.T, result *Staff, err error) {
				assert.Nil(t, result)
				assert.Error(t, err)
				appErr, ok := err.(*errors.AppError)
				assert.True(t, ok)
				assert.Equal(t, errors.ErrDatabaseOperation, appErr.Code)
			},
		},
		{
			name: "Error - Repository Error on Create",
			input: &Staff{
				FirstName: "Test Staff",
				Code:      "TEST001",
				StaffType: "doctor",
			},
			setupMocks: func(repo *mockRepository, logger *mockLogger) {
				// Only expect begin log for error case
				logger.On("Info", "service : CreateStaff : begin", mock.Anything).Return()

				repo.On("GetByCode", mock.Anything, "TEST001").Return(nil, nil)
				repo.On("Create", mock.Anything, mock.MatchedBy(func(staff *Staff) bool {
					return staff.Code == "TEST001" && staff.Status == "inactive"
				})).Return(nil, errors.New(errors.ErrDatabaseOperation, "database error", nil))
			},
			checkResult: func(t *testing.T, result *Staff, err error) {
				assert.Nil(t, result)
				assert.Error(t, err)
				appErr, ok := err.(*errors.AppError)
				assert.True(t, ok)
				assert.Equal(t, errors.ErrDatabaseOperation, appErr.Code)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mocks
			repo := new(mockRepository)
			logger := new(mockLogger)

			// Setup mocks
			tt.setupMocks(repo, logger)

			// Create service instance
			service := NewStaffService(repo, logger)

			// Execute test
			result, err := service.CreateStaff(context.Background(), tt.input)

			// Check results
			tt.checkResult(t, result, err)

			// Verify mock expectations
			repo.AssertExpectations(t)
			logger.AssertExpectations(t)
		})
	}
}
