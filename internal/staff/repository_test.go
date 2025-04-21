package staff

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/null-bd/staff-service-api/internal/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type RepositoryTestSuite struct {
	suite.Suite
	tc   *testutil.TestContainer
	repo IStaffRepository
}

func TestRepositorySuite(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}
	suite.Run(t, new(RepositoryTestSuite))
}

func (s *RepositoryTestSuite) SetupSuite() {
	ctx := context.Background()

	tc, err := testutil.SetupTestContainer(ctx)
	require.NoError(s.T(), err)
	s.tc = tc

	err = s.createSchema(ctx)
	require.NoError(s.T(), err)

	mockLogger := new(mockLogger)
	mockLogger.On("Debug", mock.Anything, mock.Anything).Return()
	s.repo = NewStaffRepository(s.tc.Pool, mockLogger)
}

func (s *RepositoryTestSuite) TearDownSuite() {
	if s.tc != nil {
		s.tc.Teardown(context.Background())
	}
}

func (s *RepositoryTestSuite) SetupTest() {
	ctx := context.Background()
	_, err := s.tc.Pool.Exec(ctx, "DELETE FROM staffs")
	require.NoError(s.T(), err)
}

func (s *RepositoryTestSuite) createSchema(ctx context.Context) error {
	schema := `
        CREATE TYPE staff_type AS ENUM ('doctor', 'nurse', 'technician', 'administrative', 'support');
		CREATE TYPE staff_status AS ENUM ('active', 'inactive', 'on-leave', 'terminated');
		
		CREATE TABLE staffs (
			id UUID PRIMARY KEY,
			branch_id UUID NOT NULL,
			organization_id UUID NOT NULL,
			first_name VARCHAR(100) NOT NULL,
			last_name VARCHAR(100) NOT NULL,
			code VARCHAR(50) NOT NULL UNIQUE,
			status staff_status NOT NULL,
			type staff_type NOT NULL,
			specialties TEXT[],  
			departments_Id UUID NOT NULL,
			departments_role TEXT[],
			departments_isprimary BOOLEAN,   
			schedule_type VARCHAR(50) NOT NULL,
			schedule_shifts TEXT[],     
			email VARCHAR(255) UNIQUE NOT NULL,
			phone VARCHAR(20) NOT NULL,
			date_of_birth DATE NOT NULL,
			gender VARCHAR(50) NOT NULL,
			address_street VARCHAR(100) NOT NULL, 
			address_city VARCHAR(50) NOT NULL,
			address_state VARCHAR(50) NOT NULL,
			address_country VARCHAR(50) NOT NULL,
			address_zipcode VARCHAR(50) NOT NULL,      
			metadata JSONB DEFAULT '{}'::JSONB,      
			created_at TIMESTAMP WITH TIME ZONE NOT NULL,
			updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
			deleted_at TIMESTAMP WITH TIME ZONE,
			UNIQUE(branch_id, code)
		);

		CREATE TYPE user_status AS ENUM ('active', 'inactive', 'blocked');
		
		CREATE TABLE users (
			id UUID PRIMARY KEY,
			branch_id UUID NOT NULL,
			organization_id UUID NOT NULL,
			first_name VARCHAR(100) NOT NULL,
			last_name VARCHAR(100) NOT NULL,
			code VARCHAR(50) NOT NULL UNIQUE,
			status user_status NOT NULL,
			user_type UUID NOT NULL,
			email VARCHAR(255) UNIQUE NOT NULL,
			phone VARCHAR(20) NOT NULL,
			date_of_birth DATE NOT NULL,
			gender VARCHAR(10) NOT NULL,
			address TEXT NOT NULL,  
			emergencycontact_name VARCHAR(100) NOT NULL,
			emergencycontact_phone VARCHAR(20) NOT NULL,
			emergencycontact_relationship VARCHAR(100) NOT NULL,  
			metadata JSONB DEFAULT '{}'::JSONB,      
			created_at TIMESTAMP WITH TIME ZONE NOT NULL,
			updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
			deleted_at TIMESTAMP WITH TIME ZONE,
			UNIQUE(branch_id, code)
		);

		CREATE TABLE usertypes (
			id UUID PRIMARY KEY,
			name VARCHAR(100) NOT NULL,
			description TEXT,
			metadata JSONB DEFAULT '{}'::JSONB,      
			created_at TIMESTAMP WITH TIME ZONE NOT NULL,
			updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
			deleted_at TIMESTAMP WITH TIME ZONE
		);

		CREATE INDEX staffs_branch_id_idx ON staffs(branch_id);
		CREATE INDEX staffs_organization_id_idx ON staffs(organization_id);
		CREATE INDEX users_branch_id_idx ON users(branch_id);
		CREATE INDEX users_organization_id_idx ON users(organization_id);
		CREATE INDEX usertypes_id_idx ON usertypes(id);
    `
	_, err := s.tc.Pool.Exec(ctx, schema)
	return err
}

// func stringPtr(s string) *string {
// 	return &s
// }

func (s *RepositoryTestSuite) TestCreate() {

	//Arrange
	ctx := context.Background()
	now := time.Now().UTC()

	staff := &Staff{
		ID:             uuid.New().String(),
		BranchID:       uuid.New().String(),
		OrganizationID: uuid.New().String(),
		FirstName:      "Test FirstName",
		LastName:       "Test LastName",
		Code:           "TEST001",
		StaffType:      "doctor",
		Status:         "inactive",
		Specialities:   []string{"speciality1", "speciality2"},
		Departments: Departments{
			DepartmentID: uuid.New().String(),
			Role:         []string{"role1", "role2"},
			IsPrimary:    true,
		},
		Schedule: Schedule{
			Type:   "type",
			Shifts: []string{"shift1", "shift2"},
		},
		Email:       "staff@hospital.com",
		Phone:       "1234567890",
		DateOfBirth: "2000-01-01",
		Gender:      "Test gender",
		Address: Address{
			Street:  "Test street",
			City:    "Test city",
			State:   "Test state",
			Country: "Test country",
			ZipCode: "123456",
		},
		CreatedAt: now.Format(time.RFC3339),
		UpdatedAt: now.Format(time.RFC3339),
	}

	// Act
	result, err := s.repo.Create(ctx, staff)

	// Assert
	assert.NoError(s.T(), err)
	assert.NotEmpty(s.T(), result.CreatedAt)
	assert.NotEmpty(s.T(), result.UpdatedAt)

	// Verify in database
	var count int
	err = s.tc.Pool.QueryRow(ctx, "SELECT COUNT(*) FROM staffs WHERE id = $1", staff.ID).Scan(&count)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), 1, count)
	_, err = s.repo.Create(ctx, staff)
	assert.Error(s.T(), err)
}
